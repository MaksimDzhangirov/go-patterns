package fileSystem

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/prototype/code/interfaces"
)

type folder struct {
	children []interfaces.Node
	name     string
}

func NewFolder(children []interfaces.Node, name string) *folder {
	return &folder{
		children: children,
		name:     name,
	}
}

func (f *folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + indentation)
	}
}

func (f *folder) Clone() interfaces.Node {
	cloneFolder := &folder{
		name: f.name + "_clone",
	}
	var tempChildren []interfaces.Node
	for _, i := range f.children {
		nodeCopy := i.Clone()
		tempChildren = append(tempChildren, nodeCopy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
