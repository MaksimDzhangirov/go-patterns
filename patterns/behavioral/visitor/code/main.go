package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/shape"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/visitor"
)

func main() {
	square := shape.NewSquare(2)
	circle := shape.NewCircle(3)
	triangle := shape.NewTriangle(1, 2, 3)

	areaCalculator := visitor.NewAreaCalculator()
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	triangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := visitor.NewMiddleCoordinates()
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	triangle.Accept(middleCoordinates)
}
