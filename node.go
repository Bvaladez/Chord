package main

type Node struct {
	Bucket  map[string]string
	Address string
}

type handlerFunc func(*Node)

// Server type that is write only
type Handler chan<- handlerFunc

func startNodeAccessor(address string) (*Node, Handler) {
	// Create channel able to access object data, channel recieves handler funcions that take obj as param
	ch := make(chan handlerFunc)
	Bucket := make(map[string]string)
	node := new(Node)
	node.Bucket = Bucket
	node.Address = address
	// Go through each handler function in the channel and give it temporary access to the state of a node
	go func() {
		for f := range ch {
			f(node)
		}
	}()
	return node, ch
}
