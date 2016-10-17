package server

type Options struct {
	Port        string
	Host        string
	Secure      bool
	Quiet       bool
	NoBrowser   bool
	Only        []string
	Ignore      string
	PathIndex   string
	Proxy       bool
	ProxyTarget string
	ProxyWhen   string
	Root        string
	Watch       bool
	Static      []string
	URL         string
}

func (o *Options) Mount() {
	if o.Secure {
		o.URL = "https://" + o.Host
	} else {
		o.URL = "http://" + o.Host
	}

	o.URL += o.Port
}

func NewOptions() *Options {
	return &Options{
		Port:        ":1307",
		Host:        "127.0.0.1",
		Secure:      false,
		Quiet:       false,
		NoBrowser:   false,
		Only:        []string{},
		Ignore:      "",
		PathIndex:   "",
		Proxy:       false,
		ProxyTarget: "",
		ProxyWhen:   "",
		Root:        "",
		Watch:       true,
		Static:      []string{},
		URL:         "",
	}
}
