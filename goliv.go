package main

import (
	"fmt"

	lib "github.com/ericmdantas/goliv/lib"
)

func main() {
	opts := lib.GetOptions()
	server := lib.NewServer(opts)

	err := server.Start()

	if err != nil {
		fmt.Printf("\nOops, something went wrong with the server. Here's the error: %s", err.Error())
	}
}
