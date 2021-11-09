package fileSystem

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/composite/code/interfaces"
)

type folder struct {
	components []interfaces.Component
	name       string
}

func NewFolder(name string) *folder {
	return &folder{
		name: name,
	}
}

func (f *folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *folder) Add(c interfaces.Component) {
	f.components = append(f.components, c)
}
