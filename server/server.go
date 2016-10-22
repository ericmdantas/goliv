package server

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func Start(opt *Options) error {
	opt.Mount()

	if err := startServer(opt); err != nil {
		return err
	}

	if err := OpenBrowser(opt); err != nil {
		return err
	}

	return nil
}

func startServer(opt *Options) error {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		indexHTMLStr, err := InjectScript(opt)

		if err != nil {
			panic(err)
		}

		return c.HTML(http.StatusOK, indexHTMLStr)
	})

	e.GET("/ws", standard.WrapHandler(handleWSConnection(opt)))

	return e.Run(standard.New(opt.Port))
}

func handleWSConnection(opt *Options) websocket.Handler {
	notifyChange := func(conn *websocket.Conn) func() {
		return func() {
			conn.Write([]byte("reload"))
		}
	}

	return websocket.Handler(func(conn *websocket.Conn) {
		cw := NewContentWatcher(opt)

		if err := cw.Watch(notifyChange(conn)); err != nil {
			panic(err)
		}
	})
}
