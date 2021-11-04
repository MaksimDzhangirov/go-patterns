package main

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/common"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/interfaces"
)

type director struct {
	builder interfaces.Builder
}

func NewDirector(b interfaces.Builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b interfaces.Builder) {
	d.builder = b
}

func (d *director) BuildHouse() common.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
