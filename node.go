package main

type Node struct {
	Messages []string
}

type handlerFunc func(*Node)

// Server type that is write only
type handler chan<- handlerFunc

func startNodeAccessor() handler {
	// Create channel able to access object data, channel recieves handler funcions that take obj as param
	ch := make(chan handlerFunc)
	state := new(Node)
	// Go through each handler function in the channel and give it temporary access to the state of a node
	go func() {
		for f := range ch {
			f(state)
		}
	}()
	return ch
}
