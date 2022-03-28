package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func shell(address string) {
	log.Printf("Starting interactive shell")
	log.Printf("Node Address: %s", address)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		args := strings.Split(line, " ")

		if len(args) > 1 {
			for i := range args {
				args[i] = strings.TrimSpace(args[i])
			}
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
			cmdGet(args)
		case "post":
			cmdPost(args)
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
		PORT = args[1]
	} else {
		log.Println("Specify port number\nUsage: port 3410")
	}
}

func cmdCreate(address string) {
	node, accessor := startNodeAccessor(address)
	go node.serve(accessor)
	log.Printf("Now listening on %s", address)
}

func cmdGet(args []string) {
	if len(args) != 3 {
		log.Println("Must specify key in get call.")
	}
		var key = args[1]
		var address = args[2]
		var value string
		if err := rpcCall(address, "Handler.Get", key, &value); err != nil {
			log.Fatalf("calling Feed.Get: %v", err)
		}
		log.Printf("%s", value)
}

func cmdPost(args []string) {
	if len(args) != 4 {
		log.Printf("%d", len(args))
		log.Printf("You must specify a key value and address in post call.")
	}
	var junk Nothing
	Post := new(KVPost)
	Post.Key = args[1]
	Post.Value = args[2]
	Post.ToAddress = args[3]

	if err := rpcCall(Post.ToAddress, "Handler.Post", Post, &junk); err != nil {
		log.Fatalf("calling Server.Post: %v", err)
	}

}
