package tv

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/command/code/interfaces"

type button struct {
	command interfaces.Command
}

func NewButton(command interfaces.Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}
