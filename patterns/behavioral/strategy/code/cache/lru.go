package cache

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/strategy/code/interfaces"
)

type lru struct {
}

func NewLru() *lru {
	return &lru{}
}

func (l *lru) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lru strategy")
}
