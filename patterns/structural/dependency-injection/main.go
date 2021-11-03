package main

import "log"

type Database interface {
	GetUser() string
}

type DefaultDatabase struct {
}

func (db DefaultDatabase) GetUser() string {
	return "user 1"
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct {
}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s! Nice to meet you", userName)
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	db := DefaultDatabase{}
	greeter := NiceGreeter{}
	program := Program{
		Db:      db,
		Greeter: greeter,
	}

	program.Execute()
}
