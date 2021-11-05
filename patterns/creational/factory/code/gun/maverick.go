package gun

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/factory/code/interfaces"

type maverick struct {
	gun
}

func NewMaverick() interfaces.Gun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
