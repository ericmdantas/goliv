package main

import (
	"log"

	lib "github.com/ericmdantas/goliv/lib"
)

func main() {
	opts := lib.GetOptions()
	server := lib.NewServer(opts)

	if err := server.Start(); err != nil {
		log.Fatalf("\nOops, something went wrong with the server. Here's the error: %s", err.Error())
	}
}
