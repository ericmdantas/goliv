package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	cfgFileName = ".golivrc"
)

type Options struct {
	Port        string   `json:"port,omitempty"`
	Host        string   `json:"host,omitempty"`
	Secure      bool     `json:"secure,omitempty"`
	Quiet       bool     `json:"quiet,omitempty"`
	NoBrowser   bool     `json:"noBrowser,omitempty"`
	Only        []string `json:"only,omitempty"`
	Ignore      string   `json:"ignore,omitempty"`
	PathIndex   string   `json:"pathIndex,omitempty"`
	Proxy       bool     `json:"proxy,omitempty"`
	ProxyWhen   string   `json:"proxyWhen,omitempty"`
	ProxyTarget string   `json:"proxyTarget,omitempty"`
	Root        string   `json:"root,omitempty"`
	Static      string   `json:"static,omitempty"`

	OnlyCLI string
	HTTPURL string
	WSURL   string
}

func (o *Options) Assign(defaultOpt, fileOpt, cliOpt Options) error {
	bDefaultValuesOpt, err := json.Marshal(defaultOpt)

	if err != nil {
		return err
	}

	bFileOpt, err := json.Marshal(fileOpt)

	if err != nil {
		return err
	}

	bCliOpt, err := json.Marshal(cliOpt)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bDefaultValuesOpt, o); err != nil {
		return err
	}

	if err := json.Unmarshal(bFileOpt, o); err != nil {
		return err
	}

	if err := json.Unmarshal(bCliOpt, o); err != nil {
		return err
	}

	if len(fileOpt.Only) != 0 {
		o.Only = fileOpt.Only
	}

	if cliOpt.OnlyCLI != "" {
		o.Only = strings.Split(o.OnlyCLI, ",")
	}

	return nil
}

func (o *Options) Parse() {
	host := "127.0.0.1"
	port := ":1308"

	if o.Host == "" {
		o.Host = host
	}

	if o.Port == "" {
		o.Port = port
	}

	if o.Secure {
		o.HTTPURL = "https://" + o.Host
		o.WSURL = "wss://" + o.Host
	} else {
		o.HTTPURL = "http://" + o.Host
		o.WSURL = "ws://" + o.Host
	}

	if len(o.Only) == 0 && o.OnlyCLI != "" {
		o.Only = strings.Split(o.OnlyCLI, ",")
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
		OnlyCLI:     "",
		Only:        []string{},
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
	pathGolivRc := filepath.Join(opt.Root, cfgFileName)

	if _, err := os.Stat(pathGolivRc); os.IsNotExist(err) {
		return Options{}, nil
	}

	info, err := ioutil.ReadFile(pathGolivRc)

	if err != nil {
		return Options{}, err
	}

	if err := json.Unmarshal(info, &opt); err != nil {
		return Options{}, err
	}

	return opt, nil
}
