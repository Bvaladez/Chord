package main

func (handler Handler) Ping(null Nothing, reply *string) error {
	finished := make(chan struct{})
	handler <- func(n *Node) {
		*reply = "pong"
	}
	<-finished
	return nil
}

func (handler Handler) Post(KV *KVPost, reply *Nothing) error {
	finished := make(chan struct{})
	// Load function into server (Actor) to queue function call and access to state changes
	handler <- func(node *Node) {
		node.Bucket[KV.Key] = KV.Value
		finished <- struct{}{}
	}
	<-finished
	return nil
}

func (handler Handler) Get(key string, reply *string) error {
	finished := make(chan struct{})
	handler <- func(node *Node) {
		if val, ok := node.Bucket[key]; ok {
			*reply = val
		}else{
			*reply = "No matching key"
		}
		finished <- struct{}{}
	}
	<-finished
	return nil
}
