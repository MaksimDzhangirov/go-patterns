package shape

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/interfaces"

type square struct {
	side int
}

func NewSquare(side int) *square {
	return &square{
		side: side,
	}
}

func (s *square) Accept(v interfaces.Visitor) {
	v.VisitForSquare(s)
}

func (s *square) GetType() string {
	return "Square"
}
