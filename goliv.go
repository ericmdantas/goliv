package main

import (
	"fmt"

	lib "github.com/ericmdantas/goliv/lib"
)

func main() {
	f := lib.IndexFile{}
	f.ReadIndex()

	fmt.Println(f.IndexHTML)
}
