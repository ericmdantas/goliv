package server

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func InjectScript(o *Options) error {
	file, err := ioutil.ReadFile(o.PathIndex + "/index.html")

	if err != nil {
		return err
	}

	fileStr := string(file)
	fileStr = strings.Replace(fileStr, "</body>", WSScript+"</body>", -1)

	fmt.Println(fileStr)

	return nil
}
