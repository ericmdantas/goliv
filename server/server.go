package server

import (
	"compress/gzip"
	"log"
	"net/http"
	_ "net/http/httputil"
	"path/filepath"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

const (
	reloadEvent = "reload"
)

func Start(cfg *Config) error {
	cliOpt := *cfg
	defaultOpt := *NewConfig()
	fileOpt, err := parseGolivRc(*cfg)

	if err != nil {
		return err
	}

	cfg.Assign(defaultOpt, fileOpt, cliOpt)
	cfg.Parse()

	if err := startServer(cfg); err != nil {
		return err
	}

	if err := openBrowser(cfg); err != nil {
		return err
	}

	return nil
}

func startServer(cfg *Config) error {
	e := echo.New()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestCompression,
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  filepath.Join(cfg.Root, cfg.PathIndex),
		HTML5: true,
		Index: "_______", // serve the index by hand
	}))

	e.GET("/", sendIndex(cfg))
	e.GET("/ws", standard.WrapHandler(handleWSConnection(cfg)))

	/*if cfg.Proxy {
		e.Get(cfg.ProxyWhen, func(c echo.Context) error {
			res := c.Response()
			req := c.Request()
			url := URL.Parse(cfg.ProxyTarget)

			return httputil.NewSingleHostReverseProxy(url).ServeHTTP(res, req)
		})
	}*/

	log.Printf("Goliv running on %s\n", cfg.HTTPURL)

	if cfg.Secure {
		return e.Run(standard.WithConfig(engine.Config{
			Address:     cfg.Port,
			TLSCertFile: "server/crt/server.crt",
			TLSKeyFile:  "server/crt/server.key",
		}))
	}

	return e.Run(standard.New(cfg.Port))
}

func sendIndex(cfg *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		f := newIndexFile(cfg)

		if err := cfg.readIndexHTML(f); err != nil {
			panic(err)
		}

		indexHTMLStr, err := injectScript(cfg)

		if err != nil {
			panic(err)
		}

		return c.HTML(http.StatusOK, indexHTMLStr)
	}
}

func handleWSConnection(cfg *Config) websocket.Handler {
	notifyChange := func(conn *websocket.Conn) func() {
		return func() {
			conn.Write([]byte(reloadEvent))
		}
	}

	return websocket.Handler(func(conn *websocket.Conn) {
		if err := watchContent(cfg, notifyChange(conn)); err != nil {
			panic(err)
		}
	})
}
