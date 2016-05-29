package goliv

import (
	"fmt"
	"io/ioutil"
)

const (
	mainHtml = "index.html"
)

func Yo() {
	fmt.Println("yo!")
}

func Read() string {
	str, err := ioutil.ReadFile(mainHtml)

	if err != nil {
		panic(err)
	}

	return string(str)
}
