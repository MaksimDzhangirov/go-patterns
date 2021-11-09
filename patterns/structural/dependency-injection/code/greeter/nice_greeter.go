package greeter

import "log"

type niceGreeter struct {
}

func NewNiceGreeter() *niceGreeter {
	return &niceGreeter{}
}

func (ng *niceGreeter) Greet(userName string) {
	log.Printf("Hello %s! Nice to meet you", userName)
}
