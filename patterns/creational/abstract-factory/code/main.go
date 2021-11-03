package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/adidas"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/nike"
)

func getSportsFactory(brand string) (interfaces.SportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Company{}, nil
	}
	if brand == "nike" {
		return &nike.Company{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s interfaces.Shoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetails(s interfaces.Short) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
