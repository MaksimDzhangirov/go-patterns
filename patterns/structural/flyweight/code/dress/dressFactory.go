package dress

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/flyweight/code/interfaces"
)

const (
	// TerroristDressType константа для одежды террориста
	TerroristDressType = "tDress"
	// CounterTerroristDressType константа для одежды спецназовца
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]interfaces.Dress),
	}
)

type dressFactory struct {
	dressMap map[string]interfaces.Dress
}

func (d *dressFactory) GetDressByType(dressType string) (interfaces.Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = NewTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = NewCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}