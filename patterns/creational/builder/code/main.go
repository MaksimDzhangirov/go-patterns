package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/igloo"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/normal"
)

func getBuilder(builderType string) interfaces.Builder {
	if builderType == "normal" {
		return normal.NewNormalBuilder()
	}
	if builderType == "igloo" {
		return igloo.NewIglooBuilder()
	}
	return nil
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetFloor())
}
