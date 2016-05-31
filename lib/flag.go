package goliv

import "flag"

func GetOptions() *Options {
	opt := NewOptions()

	flag.BoolVar(&opt.Quiet, "quiet", false, "-quiet, for no logging")
	flag.BoolVar(&opt.NoBrowser, "noBrowser", false, "-noBrowser, if you want to open the browser or not")
	flag.BoolVar(&opt.Secure, "secure", false, "-secure, if you want to work with https or not")
	flag.BoolVar(&opt.Proxy, "proxy", false, "-proxy, if you want to work with proxy or not")
	flag.StringVar(&opt.ProxyWhen, "proxyWhen", "", "-proxyWhen, when to proxy: /api/, for example")
	flag.StringVar(&opt.ProxyTarget, "proxyTarget", "", "-proxyTarget, server that'll receive the request and respond to it")
	flag.StringVar(&opt.PathIndex, "pathIndex", ".", "-pathIndex, path to index.html, default is root")
	flag.StringVar(&opt.Port, "port", "1307", "-port, port number")
	flag.StringVar(&opt.Host, "host", "127.0.0.1", "-host, host name")
	flag.StringVar(&opt.Ignore, "ignore", ".", "-ignore, regex of folders/files to ignore")
	flag.StringVar(&opt.Only, "only", ".", "-only, only folders/files to keep an eye on")

	flag.Parse()

	return opt
}
