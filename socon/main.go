package main

import (
	"log"
	"net/rpc"
	"os"
)

func main() {
	var addr string
	if len(os.Args) < 2 {
		addr = "localhost:9900"
	} else {
		addr = os.Args[1]
	}

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}

	var reply string
	client.Call("Overseer.Top", nil, &reply)
	log.Print(reply)
}
