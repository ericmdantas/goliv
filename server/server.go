package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Start(opt *Options) error {
	opt.Mount()

	startServer(opt)
	OpenBrowser(opt)

	if err := InjectScript(opt); err != nil {
		panic(err)
	}

	if err := StartWatcher(opt); err != nil {
		return err
	}

	return nil
}

func startServer(opt *Options) error {
	e := echo.New()

	e.Use(middleware.Static(""))
	e.Use(middleware.Static("/" + opt.Only))

	e.Run(standard.New(opt.Port))

	return nil
}
