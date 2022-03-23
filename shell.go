package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func shell(address string) {
	log.Printf("Starting interactive shell")
	log.Printf("Commands are: get, post")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		args := strings.SplitN(line, " ", 2)

		if len(args) > 1 {
			args[1] = strings.TrimSpace(args[1])
		} else if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			cmdHelp()
		case "quit":
			cmdQuit()
		case "port":
			cmdPort(args)
		case "create":
			cmdCreate(address)
		case "get":
			n := 10
			if len(args) == 2 {
				var err error
				if n, err = strconv.Atoi(args[1]); err != nil {
					log.Fatalf("parsing number of messages: %v", err)
				}
			}
			var messages []string
			if err := rpcCall(address, "Server.Get", n, &messages); err != nil {
				log.Fatalf("calling Feed.Get: %v", err)
			}
			for _, elt := range messages {
				log.Println(elt)
			}

		case "post":
			if len(args) != 2 {
				log.Printf("You must specify a message to post")
				continue
			}
			var junk Nothing
			if err := rpcCall(address, "Server.Post", args[1], &junk); err != nil {
				log.Fatalf("calling Server.Post: %v", err)
			}
		default:
			log.Printf("Command not recognized. Type help for a list of commands.")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}
}

func cmdHelp() {
	log.Println("Commands: help, quit, port")
}

// TODO send data to immediate succesor on quit
func cmdQuit() {
	log.Println("Shutting Down node")
	os.Exit(0)
}

func cmdPort(args []string) {
	if len(args) > 1 {
		port = args[1]
	} else {
		log.Printf("You must specify a port number to set. %s is default port", defaultPort)
	}
}

func cmdCreate(address string) {
	server(address)
}
