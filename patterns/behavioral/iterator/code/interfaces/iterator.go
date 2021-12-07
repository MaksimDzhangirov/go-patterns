package interfaces

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/model"
)

type Iterator interface {
	HasNext() bool
	GetNext() *model.User
}
