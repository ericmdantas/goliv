package main

import (
	"flag"
	"log"

	"github.com/ericmdantas/goliv/server"
)

func main() {
	o := server.NewOptions()

	flag.StringVar(&o.Port, "port", "", "a string beginning with")
	flag.StringVar(&o.Host, "host", "", "the base of your HTTPURL")
	flag.StringVar(&o.Root, "root", "", "the root of you app")
	flag.BoolVar(&o.Secure, "secure", false, "secure server or not")
	flag.BoolVar(&o.Quiet, "quiet", false, "no log")
	flag.BoolVar(&o.NoBrowser, "noBrowser", false, "doesn't open the browser")
	flag.StringVar(&o.Only, "only", "", "watchable paths - separated by comma")
	flag.StringVar(&o.Ignore, "ignore", "", "paths ignored")
	flag.StringVar(&o.PathIndex, "pathIndex", "", "path to the index.html")
	flag.BoolVar(&o.Proxy, "proxy", false, "if it's working as a reverse proxy or not")
	flag.StringVar(&o.ProxyWhen, "proxyWhen", "", "when to activate the proxy")
	flag.StringVar(&o.ProxyTarget, "proxyTarget", "", "target server/endpoint")

	flag.Parse()

	if err := server.Start(o); err != nil {
		log.Fatalln(err)
	}
}
