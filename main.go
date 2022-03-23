package main

import (
	"flag"
	"log"
)

func main() {
	var isServer bool
	var isClient bool
	var address string
	flag.BoolVar(&isServer, "server", false, "start as server")
	flag.BoolVar(&isClient, "client", false, "start as client")
	flag.Parse()
	if isServer && isClient {
		log.Fatalf("cannot be both a client and a server")
	}
	if !isServer && !isClient {
		printUsage()
	}
	switch flag.NArg() {
	case 0:
		if isClient {
			address = defaultHost + ":" + defaultPort
		} else {
			address = ":" + defaultPort
		}
	case 1:
		// User specified the address
		address = flag.Arg(0)
	default:
		printUsage()
	}
	if isClient {
		shell(address)
	} else {
		server(address)
	}
}
