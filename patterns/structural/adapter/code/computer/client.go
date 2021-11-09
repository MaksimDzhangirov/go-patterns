package computer

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/adapter/code/interfaces"

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) InsertSquareUsbInComputer(com interfaces.Computer) {
	com.InsertInSquarePort()
}
