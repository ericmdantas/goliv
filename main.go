package main

import (
	"flag"
	"log"

	"github.com/ericmdantas/goliv/server"
)

func main() {
	cfg := server.NewConfig()

	flag.StringVar(&cfg.Port, "port", "", "the port number begining with :")
	flag.StringVar(&cfg.Host, "host", "", "the base of your HTTPURL")
	flag.StringVar(&cfg.Root, "root", "", "the root of your project")
	flag.BoolVar(&cfg.HTTP2, "http2", false, "http2 server or not")
	flag.BoolVar(&cfg.Quiet, "quiet", false, "no log")
	flag.BoolVar(&cfg.OpenBrowser, "openBrowser", false, "opens the browser")
	flag.StringVar(&cfg.OnlyCLI, "only", "", "watchable paths - separated by comma")
	flag.StringVar(&cfg.IgnoreCLI, "ignore", "", "paths ignored - separated by comma")
	flag.StringVar(&cfg.PathIndex, "pathIndex", "", "path to the index.html")
	flag.BoolVar(&cfg.Proxy, "proxy", false, "reverse proxy")
	flag.StringVar(&cfg.ProxyWhen, "proxyWhen", "", "when to activate the proxy")
	flag.StringVar(&cfg.ProxyTarget, "proxyTarget", "", "target server/endpoint")

	flag.Parse()

	if err := server.Start(cfg); err != nil {
		log.Fatalln(err)
	}
}
