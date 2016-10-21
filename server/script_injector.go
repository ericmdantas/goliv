package server

import (
	"io/ioutil"
	"strings"
)

func InjectScript(o *Options) (string, error) {
	file, err := ioutil.ReadFile(o.PathIndex + "/index.html")

	if err != nil {
		return "", err
	}

	fileWithScript := strings.Replace(string(file), "</body>", WSScript+"</body>", -1)

	return fileWithScript, nil
}
