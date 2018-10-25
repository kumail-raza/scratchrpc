package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	err = client.Call("net.rpc.SayHello", "hey", &reply)
	if err != nil {
		log.Fatal("Couldnt call client")
	}

	fmt.Println("REPLY", reply)
}
