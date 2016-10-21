package server

import (
	"net/http"

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

	if err := StartWatcher(opt); err != nil {
		return err
	}

	return nil
}

func startServer(opt *Options) error {
	e := echo.New()

	indexHTMLStr, err := InjectScript(opt)

	if err != nil {
		return err
	}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, indexHTMLStr)
	})

	return e.Run(standard.New(opt.Port))
}
