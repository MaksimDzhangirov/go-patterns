package igloo

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/common"

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) GetHouse() common.House {
	return common.NewHouse(b.windowType, b.doorType, b.floor)
}
