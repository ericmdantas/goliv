package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	cfgFileName         = ".golivrc"
	defaultHost         = "127.0.0.1"
	defaultPort         = ":1308"
	inlinePathSeparator = ","
	nameIndexHTMLFile   = "index.html"
)

type Config struct {
	Port        string   `json:"port,omitempty"`
	Host        string   `json:"host,omitempty"`
	Secure      bool     `json:"secure,omitempty"`
	Quiet       bool     `json:"quiet,omitempty"`
	NoBrowser   bool     `json:"noBrowser,omitempty"`
	Only        []string `json:"only,omitempty"`
	Ignore      []string `json:"ignore,omitempty"`
	PathIndex   string   `json:"pathIndex,omitempty"`
	Proxy       bool     `json:"proxy,omitempty"`
	ProxyWhen   string   `json:"proxyWhen,omitempty"`
	ProxyTarget string   `json:"proxyTarget,omitempty"`
	Root        string   `json:"root,omitempty"`
	Static      string   `json:"static,omitempty"`

	IgnoreCLI string
	OnlyCLI   string
	HTTPURL   string
	WSURL     string

	indexHTMLPath    string
	indexHTMLContent []byte
}

func (cfg *Config) assign(defaultOpt, fileOpt, cliOpt Config) error {
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

	if err := json.Unmarshal(bDefaultValuesOpt, cfg); err != nil {
		return err
	}

	if err := json.Unmarshal(bFileOpt, cfg); err != nil {
		return err
	}

	if err := json.Unmarshal(bCliOpt, cfg); err != nil {
		return err
	}

	if len(fileOpt.Only) != 0 {
		cfg.Only = fileOpt.Only
	}

	if cliOpt.OnlyCLI != "" {
		cfg.Only = strings.Split(cfg.OnlyCLI, inlinePathSeparator)
	}

	if len(fileOpt.Ignore) != 0 {
		cfg.Ignore = fileOpt.Ignore
	}

	if cliOpt.IgnoreCLI != "" {
		cfg.Ignore = strings.Split(cfg.IgnoreCLI, inlinePathSeparator)
	}

	return nil
}

func (cfg *Config) parse() {
	if cfg.Host == "" {
		cfg.Host = defaultHost
	}

	if cfg.Port == "" {
		cfg.Port = defaultPort
	}

	if cfg.Secure {
		cfg.HTTPURL = "https://" + cfg.Host
		cfg.WSURL = "wss://" + cfg.Host
	} else {
		cfg.HTTPURL = "http://" + cfg.Host
		cfg.WSURL = "ws://" + cfg.Host
	}

	if len(cfg.Only) == 0 {
		if cfg.OnlyCLI == "" {
			if cfg.Root == "" {
				cfg.Only = []string{"."}
			} else {
				cfg.Only = []string{cfg.Root}
			}
		} else {
			pathsSplit := strings.Split(cfg.OnlyCLI, inlinePathSeparator)

			for _, v := range pathsSplit {
				cfg.Only = append(cfg.Only, filepath.Join(cfg.Root, v))
			}
		}
	} else {
		for i := range cfg.Only {
			cfg.Only[i] = filepath.Join(cfg.Root, cfg.Only[i])
		}
	}

	for i, v := range cfg.Only {
		if str := strings.Replace(v, "*", "", -1); str != v {
			for strings.HasSuffix(str, "\\") || strings.HasSuffix(str, "/") {
				str = str[:len(str)-1]
			}

			if str == "" {
				str = filepath.Join(cfg.Root, ".")
			}

			cfg.Only[i] = str
		}
	}

	if cfg.IgnoreCLI != "" {
		cfg.Ignore = strings.Split(cfg.IgnoreCLI, inlinePathSeparator)
	}

	for i := range cfg.Ignore {
		cfg.Ignore[i] = filepath.Join(cfg.Root, cfg.Ignore[i])
	}

	cfg.HTTPURL += cfg.Port
	cfg.WSURL += cfg.Port + "/ws"

	cfg.indexHTMLPath = filepath.Join(cfg.Root, cfg.PathIndex, nameIndexHTMLFile)
}

func (cfg *Config) readIndexHTML(f IndexFileReader) error {
	indexHTMLInfo, err := f.readIndexHTML()

	if err != nil {
		return err
	}

	cfg.indexHTMLContent = indexHTMLInfo

	return nil
}

func NewConfig() *Config {
	return &Config{
		Port:        defaultPort,
		Host:        defaultHost,
		Secure:      false,
		Quiet:       false,
		NoBrowser:   false,
		OnlyCLI:     "",
		Only:        []string{},
		Ignore:      []string{},
		PathIndex:   "",
		Proxy:       false,
		ProxyTarget: "",
		ProxyWhen:   "",
		Root:        "",
		Static:      "",

		HTTPURL: "",
		WSURL:   "",

		indexHTMLPath:    "",
		indexHTMLContent: []byte{},
	}
}

func parseGolivRc(cfg Config) (Config, error) {
	pathGolivRc := filepath.Join(cfg.Root, cfgFileName)

	if _, err := os.Stat(pathGolivRc); os.IsNotExist(err) {
		return Config{}, nil
	}

	info, err := ioutil.ReadFile(pathGolivRc)

	if err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal(info, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
