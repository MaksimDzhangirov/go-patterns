package main

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/dependency-injection/code/db"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/dependency-injection/code/greeter"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/dependency-injection/code/program"
)

func main() {
	ddb := db.NewDefaultDatabase()
	nGreeter := greeter.NewNiceGreeter()
	prog := program.NewProgram(ddb, nGreeter)

	prog.Execute()
}
