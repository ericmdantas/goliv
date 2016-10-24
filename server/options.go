package server

type Options struct {
	Port        string
	Host        string
	Secure      bool
	Quiet       bool
	NoBrowser   bool
	Only        string
	Ignore      string
	PathIndex   string
	Proxy       bool
	ProxyTarget string
	ProxyWhen   string
	Root        string
	Static      string

	HTTPURL string
	WSURL   string
}

func (o *Options) Mount() {
	if o.Secure {
		o.HTTPURL = "https://" + o.Host
		o.WSURL = "wss://" + o.Host
	} else {
		o.HTTPURL = "http://" + o.Host
		o.WSURL = "ws://" + o.Host
	}

	o.HTTPURL += o.Port
	o.WSURL += o.Port + "/ws"
}

func (o *Options) Merge(opt ...Options) Options {
	return Options{}
}

func NewOptions() *Options {
	return &Options{
		Port:        ":1308",
		Host:        "127.0.0.1",
		Secure:      false,
		Quiet:       false,
		NoBrowser:   false,
		Only:        "",
		Ignore:      "",
		PathIndex:   "",
		Proxy:       false,
		ProxyTarget: "",
		ProxyWhen:   "",
		Root:        "",
		Static:      "",
		HTTPURL:     "",
		WSURL:       "",
	}
}
