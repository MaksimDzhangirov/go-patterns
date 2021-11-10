package player

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/flyweight/code/dress"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/flyweight/code/interfaces"
)

type player struct {
	dress      interfaces.Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *player {
	pDress, _ := dress.GetDressFactorySingleInstance().GetDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      pDress,
	}
}

func (p *player) NewLocation(lat int, long int) {
	p.lat = lat
	p.long = long
}

func (p *player) GetPlayerType() string {
	return p.playerType
}

func (p *player) GetDressColor() string {
	return p.dress.GetColor()
}
