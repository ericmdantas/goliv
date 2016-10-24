package server

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const (
	cfgFileName = ".golivrc"
)

type Options struct {
	Port        string `json:",omitempty"`
	Host        string `json:",omitempty"`
	Secure      bool   `json:",omitempty"`
	Quiet       bool   `json:",omitempty"`
	NoBrowser   bool   `json:",omitempty"`
	Only        string `json:",omitempty"`
	Ignore      string `json:",omitempty"`
	PathIndex   string `json:",omitempty"`
	Proxy       bool   `json:",omitempty"`
	ProxyTarget string `json:",omitempty"`
	ProxyWhen   string `json:",omitempty"`
	Root        string `json:",omitempty"`
	Static      string `json:",omitempty"`

	HTTPURL string
	WSURL   string
}

func (o *Options) Assign(fileOpt, cliOpt Options) error {
	bFileOpt, err := json.Marshal(fileOpt)

	if err != nil {
		return err
	}

	bCliOpt, err := json.Marshal(cliOpt)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bFileOpt, o); err != nil {
		return err
	}

	if err := json.Unmarshal(bCliOpt, o); err != nil {
		return err
	}

	return nil
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

func parseGolivRc(opt Options) (Options, error) {
	info, err := ioutil.ReadFile(filepath.Join(opt.Root, cfgFileName))

	if err != nil {
		return Options{}, err
	}

	if err := json.Unmarshal(info, &opt); err != nil {
		return Options{}, err
	}

	return opt, nil
}
