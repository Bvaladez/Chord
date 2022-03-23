package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	defaultHost = "localhost"
	defaultPort = "3410"
)

var port = defaultPort

type Nothing struct{}

func server(address string) {
	accessor := startNodeAccessor()
	rpc.Register(accessor)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("http.Server: %v", err)
	}
}

func rpcCall(address string, method string, request interface{}, response interface{}) error {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Printf("rpc.DialHTTP: %v", err)
		return err
	}
	defer client.Close()

	if err = client.Call(method, request, response); err != nil {
		log.Printf("client.Call %s: %v", method, err)
		return err
	}
	return nil
}

func (h Handler) Ping(null Nothing, reply *string) error {
	finished := make(chan struct{})
	h <- func(n *Node) {
		*reply = "pong"
	}
	<-finished
	return nil
}

func (h Handler) Post(msg string, reply *Nothing) error {
	finished := make(chan struct{})
	// Load function into server (Actor) to queue function call and access to state changes
	h <- func(f *Node) {
		f.Messages = append(f.Messages, msg)
		finished <- struct{}{}
	}
	<-finished
	return nil
}

func (h Handler) Get(count int, reply *[]string) error {
	finished := make(chan struct{})
	h <- func(f *Node) {
		if len(f.Messages) < count {
			count = len(f.Messages)
		}
		*reply = make([]string, count)
		copy(*reply, f.Messages[len(f.Messages)-count:])
		finished <- struct{}{}
	}
	<-finished
	return nil
}
