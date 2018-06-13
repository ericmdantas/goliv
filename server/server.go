package server

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"time"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
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

	cfg.parse()

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
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestCompression,
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  s.cfg.Root,
		HTML5: true,
		Index: "^';..;'^", // served by hand
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  filepath.Join(s.cfg.Root, s.cfg.PathIndex),
		HTML5: true,
		Index: "^';..;'^", // served by hand
	}))
	e.GET("/", s.sendIndex())
	e.GET("/ws", s.handleWSConnection)

	if s.cfg.Proxy {
		e.GET(s.cfg.ProxyWhen, func(c echo.Context) error {
			u, err := url.Parse(s.cfg.ProxyTarget)

			if err != nil {
				return err
			}

			httputil.NewSingleHostReverseProxy(u).ServeHTTP(c.Response(), c.Request())

			return nil
		})
	}

	e.GET("/*", s.sendIndex())

	log.Printf("Goliv running on %s\n", s.cfg.HTTPURL)

	if err := cbServerReady(); err != nil {
		return err
	}

	if s.cfg.HTTP2 {
		err := e.StartTLS(s.cfg.Port, "server/crt/server.crt", "server/crt/server.key")

		if err != nil {
			return err
		}
	}

	return e.Start(s.cfg.Port)
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

func (s *server) handleWSConnection(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		s.onChange(ws)
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (s *server) onChange(ws *websocket.Conn) {
	select {
	case event := <-s.watcher.Event:
		switch event.Op {
		case watcher.Create:
			if !s.cfg.Quiet {
				log.Println("CREATEd -> ", event.Name())
			}

			s.notifyChange(ws)
		case watcher.Write:
			if !s.cfg.Quiet {
				log.Println("CHANGED -> ", event.Name())
			}

			s.notifyChange(ws)
		case watcher.Remove:
			if !s.cfg.Quiet {
				log.Println("REMOVED -> ", event.Name())
			}

			s.notifyChange(ws)
		case watcher.Rename:
			if !s.cfg.Quiet {
				log.Println("RENAMED -> ", event.Name())
			}

			s.notifyChange(ws)
		}
	case err := <-s.watcher.Error:
		log.Fatalln(err)
	}
}

func (s *server) notifyChange(ws *websocket.Conn) {
	ws.Write([]byte(reloadEvent))
}

func (s *server) startWatcher() error {
	s.watcher = watcher.New()
	s.watcher.SetMaxEvents(1)

	for _, path := range s.cfg.Only {
		if err := s.watcher.Add(path); err != nil {
			return err
		}
	}

	for _, path := range s.cfg.Ignore {
		if err := s.watcher.Ignore(path); err != nil {
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
