package fileSystem

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/prototype/code/interfaces"
)

type file struct {
	name string
}

func NewFile(name string) *file {
	return &file{
		name: name,
	}
}

func (f *file) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) Clone() interfaces.Node {
	return &file{name: f.name + "_clone"}
}
