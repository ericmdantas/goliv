package server

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const (
	cfgFileName = ".golivrc"
)

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
