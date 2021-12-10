package cache

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/strategy/code/interfaces"
)

type fifo struct {
}

func NewFifo() *fifo {
	return &fifo{}
}

func (l *fifo) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by fifo strategy")
}
