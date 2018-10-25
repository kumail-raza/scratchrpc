package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/minhajuddinkhan/scratchrpc/services"
)

func main() {

	//	h := new(services.HelloService)
	//rpc.Register(h)
	//	h := services.HelloService{Greeting: "what"}
	rpc.RegisterName("net.rpc", new(services.HelloService))
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listener, nil)

}
