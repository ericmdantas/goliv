package server

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	_ "net/http/httputil"
	"path/filepath"
	"time"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/radovskyb/watcher"
)

const (
	reloadEvent       = "reload"
	intervalFileCheck = 1000
)

func Start(cfg *Config) error {
	defaultOpt := *NewConfig()
	fileOpt, err := parseGolivRc(*cfg)
	cliOpt := *cfg

	if err != nil {
		return err
	}

	if err := cfg.assign(defaultOpt, fileOpt, cliOpt); err != nil {
		return fmt.Errorf("There was an error when assigning the properties. %v\n", err)
	}

	cfg.Parse()

	s := &server{
		cfg: cfg,
	}

	return s.start(func() error {
		if err := s.startWatcher(); err != nil {
			return err
		}

		return openBrowser(s.cfg)
	})
}

type server struct {
	cfg     *Config
	watcher *watcher.Watcher
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

	return e.Run(standard.New(s.cfg.Port))
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
		s.onChange(conn)
	})
}

func (s *server) onChange(conn *websocket.Conn) {
	select {
	case event := <-s.watcher.Event:
		switch event.Op {
		case watcher.Create:
			if !s.cfg.Quiet {
				log.Println("Created file:", event.Name())
			}

			s.notifyChange(conn)
		case watcher.Write:
			if !s.cfg.Quiet {
				log.Println("Changed file:", event.Name())
			}

			s.notifyChange(conn)
		case watcher.Remove:
			if !s.cfg.Quiet {
				log.Println("Removed file:", event.Name())
			}

			s.notifyChange(conn)
		case watcher.Rename:
			if !s.cfg.Quiet {
				log.Println("Renamed file:", event.Name())
			}

			s.notifyChange(conn)
		}
	case err := <-s.watcher.Error:
		log.Fatalln(err)
	}
}

func (s *server) notifyChange(conn *websocket.Conn) {
	conn.Write([]byte(reloadEvent))
}

func (s *server) startWatcher() error {
	s.watcher = watcher.New()
	s.watcher.SetMaxEvents(1)

	for _, path := range s.cfg.Only {
		if err := s.watcher.Add(path); err != nil {
			return err
		}
	}

	go func() {
		if err := s.watcher.Start(time.Millisecond * intervalFileCheck); err != nil {
			panic(err)
		}
	}()

	return nil
}
