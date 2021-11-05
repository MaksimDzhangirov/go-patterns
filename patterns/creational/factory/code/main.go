package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/factory/code/gun"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/factory/code/interfaces"
	"log"
)

func main() {
	ak47, err := gun.GetGun("ak47")
	if err != nil {
		log.Fatalf("Cannot create ak47 gun. Error %v", err)
	}
	maverick, err := gun.GetGun("maverick")
	if err != nil {
		log.Fatalf("Cannot create maverick gun. Error %v", err)
	}
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g interfaces.Gun) {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}
