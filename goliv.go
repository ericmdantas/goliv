package main

import (
	"fmt"

	lib "github.com/ericmdantas/goliv/lib"
)

func main() {
	lib.Yo()
	opt := lib.GetOptionsFromFlags()

	fmt.Println(opt)
}
