package cache

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/strategy/code/interfaces"
)

type lfu struct {
}

func NewLfu() *lfu {
	return &lfu{}
}

func (l *lfu) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lfu strategy")
}
