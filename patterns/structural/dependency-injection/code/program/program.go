package program

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/dependency-injection/code/interfaces"

type program struct {
	Db      interfaces.Database
	Greeter interfaces.Greeter
}

func NewProgram(db interfaces.Database, greeter interfaces.Greeter) *program {
	return &program{
		Db:      db,
		Greeter: greeter,
	}
}

func (p *program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}
