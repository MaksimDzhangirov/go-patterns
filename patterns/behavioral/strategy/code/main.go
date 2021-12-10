package main

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/strategy/code/cache"

func main() {
	lfu := cache.NewLfu()
	cache1 := cache.InitCache(lfu)
	cache1.Add("a", "1")
	cache1.Add("b", "2")
	cache1.Add("c", "3")
	lru := cache.NewLru()
	cache1.SetEvictionAlgo(lru)
	cache1.Add("d", "4")
	fifo := cache.NewFifo()
	cache1.SetEvictionAlgo(fifo)
	cache1.Add("e", "5")
}
