package main

import (
	"log"

	"github.com/ericmdantas/goliv/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
