package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/prototype/code/fileSystem"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/prototype/code/interfaces"
)

func main() {
	file1 := fileSystem.NewFile("File1")
	file2 := fileSystem.NewFile("File2")
	file3 := fileSystem.NewFile("File3")
	folder1 := fileSystem.NewFolder([]interfaces.Node{file1}, "Folder1")
	folder2 := fileSystem.NewFolder([]interfaces.Node{file1, file2, file3}, "Folder2")
	fmt.Println("Printing hierarchy for Folder 1")
	folder1.Print("    ")
	fmt.Println("\nPrinting hierarchy for Folder 2")
	folder2.Print("    ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for Clone Folder 2")
	cloneFolder.Print("    ")
}
