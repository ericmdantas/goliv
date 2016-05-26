package goliv

import "flag"

func GetOptionsFromFlags() *Options {
	opt := NewOptions()

	noBrowser := flag.Bool("noBrowser", opt.NoBrowser, "")
	host := flag.String("host", opt.Host, "")
	secure := flag.Bool("secure", opt.Secure, "")
	port := flag.String("port", opt.Port, "")
	pathIndex := flag.String("pathIndex", opt.PathIndex, "")
	quiet := flag.Bool("quiet", opt.Quiet, "")
	proxy := flag.Bool("proxy", opt.Proxy, "")
	proxyTarget := flag.String("proxyTarget", opt.ProxyTarget, "")
	proxyWhen := flag.String("proxyWhen", opt.ProxyWhen, "")
	ignore := flag.String("ignore", opt.Ignore, "")
	only := flag.String("only", opt.Only, "")

	flag.Parse()

	opt.NoBrowser = *noBrowser
	opt.Host = *host
	opt.Secure = *secure
	opt.Port = *port
	opt.PathIndex = *pathIndex
	opt.Quiet = *quiet
	opt.Proxy = *proxy
	opt.ProxyTarget = *proxyTarget
	opt.ProxyWhen = *proxyWhen
	opt.Ignore = *ignore
	opt.Only = *only

	return opt
}
