package shape

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/visitor/code/interfaces"

type Triangle struct {
	a int
	b int
	c int
}

func NewTriangle(a, b, c int) *Triangle {
	return &Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t *Triangle) Accept(v interfaces.Visitor) {
	v.VisitForTriangle(t)
}

func (t *Triangle) GetType() string {
	return "Triangle"
}
