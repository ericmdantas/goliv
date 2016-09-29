package goliv

import "io/ioutil"

const (
	mainHTML = "index.html"
)

func InjectScriptWs(contentScript string) string {
	bs, err := ioutil.ReadFile(mainHTML)

	if err != nil {
		panic(err)
	}

	return string(bs)
}
