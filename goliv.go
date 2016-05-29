package main

import lib "github.com/ericmdantas/goliv/lib"

func main() {
	opts := lib.GetOptions()
	server := lib.NewServer(opts)
	server.Start()
}
