package shop

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/observer/code/interfaces"
)

type item struct {
	observerList []interfaces.Observer
	name         string
	inStock      bool
}

func NewItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *item) Register(o interfaces.Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) Deregister(o interfaces.Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

func removeFromSlice(observerList []interfaces.Observer, observerToRemove interfaces.Observer) []interfaces.Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
