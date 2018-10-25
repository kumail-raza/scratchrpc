package facade

import (
	"fmt"
)

//Hello Hello
type Hello struct {
	Greetings string
}

//SayHello SayHello
func (h *Hello) SayHello(name string) string {

	return fmt.Sprintf("%s , %s", h.Greetings, name)
}
