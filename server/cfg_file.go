package server

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const (
	cfgFileName = ".golivrc"
)

func parseGolivRc(opt *Options) error {
	info, err := ioutil.ReadFile(filepath.Join(opt.Root, cfgFileName))

	if err != nil {
		return nil
	}

	return json.Unmarshal(info, &opt)
}
