package main

import (
	"flag"
	"log"

	"github.com/ericmdantas/goliv/server"
)

func main() {
	o := server.NewOptions()

	flag.StringVar(&o.Port, "port", o.Port, "a string beginning with")
	flag.StringVar(&o.Host, "host", o.Host, "the base of your HTTPURL")
	flag.BoolVar(&o.Secure, "secure", o.Secure, "secure server or not")
	flag.BoolVar(&o.Quiet, "quiet", o.Quiet, "no log")
	flag.BoolVar(&o.NoBrowser, "noBrowser", o.NoBrowser, "doesn't open the browser")
	flag.StringVar(&o.Only, "only", o.Only, "watchable paths - separated by comma")
	flag.StringVar(&o.Ignore, "ignore", o.Ignore, "paths ignored")
	flag.StringVar(&o.PathIndex, "pathIndex", o.PathIndex, "path to the index.html")
	flag.BoolVar(&o.Proxy, "proxy", o.Proxy, "if it's working as a reverse proxy or not")
	flag.StringVar(&o.ProxyWhen, "proxyWhen", o.ProxyWhen, "when to activate the proxy")
	flag.StringVar(&o.ProxyTarget, "proxyTarget", o.ProxyTarget, "target server/endpoint")

	flag.Parse()

	if err := server.Start(o); err != nil {
		log.Fatalln(err)
	}
}
