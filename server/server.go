package server

import (
	"compress/gzip"
	"fmt"
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
	defaultOpt := *NewConfig()
	fileOpt, err := parseGolivRc(*cfg)
	cliOpt := *cfg

	if err != nil {
		return err
	}

	if err := cfg.Assign(defaultOpt, fileOpt, cliOpt); err != nil {
		return fmt.Errorf("There was an error when assigning the properties. %v\n", err)
	}

	cfg.Parse()

	s := server{
		cfg: cfg,
	}

	return s.start(func() error {
		return openBrowser(s.cfg)
	})
}

type server struct {
	cfg *Config
}

func (s *server) start(cbServerReady func() error) error {
	e := echo.New()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestCompression,
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  filepath.Join(s.cfg.Root, s.cfg.PathIndex),
		HTML5: true,
		Index: "_______", // serve the index by hand
	}))

	e.GET("/", s.sendIndex())
	e.GET("/ws", standard.WrapHandler(s.handleWSConnection()))

	/*if cfg.Proxy {
		e.Get(cfg.ProxyWhen, func(c echo.Context) error {
			res := c.Response()
			req := c.Request()
			url := URL.Parse(cfg.ProxyTarget)

			return httputil.NewSingleHostReverseProxy(url).ServeHTTP(res, req)
		})
	}*/

	log.Printf("Goliv running on %s\n", s.cfg.HTTPURL)

	if err := cbServerReady(); err != nil {
		return err
	}

	if s.cfg.Secure {
		err := e.Run(standard.WithConfig(engine.Config{
			Address:     s.cfg.Port,
			TLSCertFile: "server/crt/server.crt",
			TLSKeyFile:  "server/crt/server.key",
		}))

		if err != nil {
			return err
		}
	}

	if err := e.Run(standard.New(s.cfg.Port)); err != nil {
		return err
	}

	return nil
}

func (s *server) sendIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		f := newIndexFile(s.cfg)

		if err := s.cfg.readIndexHTML(f); err != nil {
			panic(err)
		}

		indexHTMLStr, err := injectScript(s.cfg)

		if err != nil {
			panic(err)
		}

		return c.HTML(http.StatusOK, indexHTMLStr)
	}
}

func (s *server) handleWSConnection() websocket.Handler {
	return websocket.Handler(func(conn *websocket.Conn) {
		if err := watchContent(s.cfg, s.notifyChange(conn)); err != nil {
			panic(err)
		}
	})
}

func (s *server) notifyChange(conn *websocket.Conn) func() {
	return func() {
		conn.Write([]byte(reloadEvent))
	}
}
