package server

import (
	"compress/gzip"
	"log"
	"net/http"
	_ "net/http/httputil"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

const reloadEvent = "reload"

func Start(opt *Options) error {
	cliOpt := *opt
	defaultOpt := *NewOptions()
	fileOpt, err := parseGolivRc(*opt)

	if err != nil {
		return err
	}

	opt.Assign(defaultOpt, fileOpt, cliOpt)
	opt.Mount()

	if err := startServer(opt); err != nil {
		return err
	}

	if err := openBrowser(opt); err != nil {
		return err
	}

	return nil
}

func startServer(opt *Options) error {
	e := echo.New()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestCompression,
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  opt.Root,
		HTML5: true,
		Index: "_______", // serve the index by hand
	}))

	e.GET("/", sendIndex(opt))
	e.GET("/ws", standard.WrapHandler(handleWSConnection(opt)))

	/*if opt.Proxy {
		e.Get(opt.ProxyWhen, func(c echo.Context) error {
			res := c.Response()
			req := c.Request()
			url := URL.Parse(opt.ProxyTarget)

			return httputil.NewSingleHostReverseProxy(url).ServeHTTP(res, req)
		})
	}*/

	log.Printf("Goliv running on %s\n", opt.HTTPURL)

	return e.Run(standard.New(opt.Port))
}

func sendIndex(opt *Options) echo.HandlerFunc {
	return func(c echo.Context) error {
		indexHTMLStr, err := injectScript(opt)

		if err != nil {
			panic(err)
		}

		return c.HTML(http.StatusOK, indexHTMLStr)
	}
}

func handleWSConnection(opt *Options) websocket.Handler {
	notifyChange := func(conn *websocket.Conn) func() {
		return func() {
			conn.Write([]byte(reloadEvent))
		}
	}

	return websocket.Handler(func(conn *websocket.Conn) {
		cw := newContentWatcher(opt)

		if err := cw.Watch(notifyChange(conn)); err != nil {
			panic(err)
		}
	})
}
