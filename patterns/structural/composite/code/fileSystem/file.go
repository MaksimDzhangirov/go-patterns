package fileSystem

import "fmt"

type file struct {
	name string
}

func NewFile(name string) *file {
	return &file{
		name: name,
	}
}

func (f *file) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) GetName() string {
	return f.name
}
