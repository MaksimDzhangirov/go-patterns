package interfaces

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/builder/code/common"

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() common.House
}