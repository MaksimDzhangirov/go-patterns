package main

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/composite/code/fileSystem"

func main() {
	file1 := fileSystem.NewFile("File1")
	file2 := fileSystem.NewFile("File2")
	file3 := fileSystem.NewFile("File3")
	folder1 := fileSystem.NewFolder("Folder1")
	folder1.Add(file1)
	folder2 := fileSystem.NewFolder("Folder2")
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)
	folder2.Search("rose")
}
