package normal

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/common"

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) SetNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) GetHouse() common.House {
	return common.NewHouse(b.windowType, b.doorType, b.floor)
}