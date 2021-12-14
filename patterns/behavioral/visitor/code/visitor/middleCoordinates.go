package visitor

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/interfaces"
)

type middleCoordinates struct {
	x int
	y int
}

func NewMiddleCoordinates() *middleCoordinates {
	return &middleCoordinates{}
}

func (mc *middleCoordinates) VisitForSquare(s interfaces.Shape) {
	// Вычисляем координаты центра масс квадрата. После вычисления присвоить их
	// в переменные x и y экземпляра.
	fmt.Println("Calculating middle point coordinates for square")
}

func (mc *middleCoordinates) VisitForCircle(s interfaces.Shape) {
	// Вычисляем координаты центра масс окружности. После вычисления присвоить их
	// в переменные x и y экземпляра.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (mc *middleCoordinates) VisitForTriangle(s interfaces.Shape) {
	// Вычисляем координаты центра масс треугольника. После вычисления присвоить их
	// в переменные x и y экземпляра.
	fmt.Println("Calculating middle point coordinates for triangle")
}
