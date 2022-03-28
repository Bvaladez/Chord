package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	DEFAULTHOST = "localhost"
)

var PORT = ""

type Nothing struct{}

type KVPost struct {
	ToAddress string
	Key   string
	Value string
}


func (node *Node) serve(accessor Handler) {
	rpc.Register(accessor)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", node.Address)
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
