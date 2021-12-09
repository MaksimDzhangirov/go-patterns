package main

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/observer/code/shop"

func main() {
	shirtItem := shop.NewItem("Nike Shirt")
	observerFirst := shop.NewCustomer("abc@gmail.com")
	observerSecond := shop.NewCustomer("xyz@gmail.com")
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}
