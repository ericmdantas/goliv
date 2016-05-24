package goliv

type Options struct {
	NoBrowser   bool
	Host        string
	Secure      bool
	Port        string
	PathIndex   string
	Quiet       bool
	Proxy       bool
	ProxyTarget string
	ProxyWhen   string
	Ignore      string
	Only        string
}

func NewOptions() *Options {
	return &Options{
		NoBrowser:   false,
		Host:        "127.0.0.1",
		Secure:      false,
		Port:        "1307",
		PathIndex:   "",
		Quiet:       false,
		Proxy:       false,
		ProxyTarget: "",
		ProxyWhen:   "",
		Ignore:      "",
		Only:        ".",
	}
}
