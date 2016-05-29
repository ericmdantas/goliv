package goliv

import "io/ioutil"

const (
	mainHTML = "index.html"
)

type IndexFile struct {
	IndexHTML string
}

func (f *IndexFile) ReadIndex() {
	bs, err := ioutil.ReadFile(mainHTML)

	if err != nil {
		panic(err)
	}

	f.IndexHTML = string(bs)
}
