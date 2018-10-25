package services

import (
	"github.com/minhajuddinkhan/scratchrpc/facade"
)

//HelloService HelloService
type HelloService struct {
	Greeting string
}

var (
	//HelloPlaceHolder HelloPlaceHolder
	HelloPlaceHolder = facade.Hello{Greetings: "Hey."}
)

//SayHello SayHello
func (h *HelloService) SayHello(in string, out *string) error {

	*out = HelloPlaceHolder.SayHello(in)
	return nil

}
