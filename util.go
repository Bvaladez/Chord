package main

import (
	"flag"
	"log"
	"os"
)

func printUsage() {
	log.Printf("Usage: %s [-server or -client] [address]", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
