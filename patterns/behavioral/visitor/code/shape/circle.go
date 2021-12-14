package shape

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/interfaces"

type circle struct {
	radius int
}

func NewCircle(radius int) *circle {
	return &circle{
		radius: radius,
	}
}

func (c *circle) Accept(v interfaces.Visitor) {
	v.VisitForCircle(c)
}

func (c *circle) GetType() string {
	return "Circle"
}
