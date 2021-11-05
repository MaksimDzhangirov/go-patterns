package gun

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/factory/code/interfaces"

type ak47 struct {
	gun
}

func NewAk47() interfaces.Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
