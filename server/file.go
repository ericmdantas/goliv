package server

import (
	"io/ioutil"
)

type IndexFileReader interface {
	readIndexHTML() ([]byte, error)
}

type indexFile struct {
	cfg *Config
}

func (f *indexFile) readIndexHTML() ([]byte, error) {
	return ioutil.ReadFile(f.cfg.indexHTMLPath)
}

func newIndexFile(cfg *Config) *indexFile {
	return &indexFile{
		cfg: cfg,
	}
}
