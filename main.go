package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var port string
	fmt.Println("Please enter port number.")
	fmt.Scanln(&port)
	address := DEFAULTHOST + ":" + port
	shell(address)
}

func old_main() {
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
			address = DEFAULTHOST + ":" + PORT
		} else {
			address = ":" + PORT
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
		//server(address)
	}
}
