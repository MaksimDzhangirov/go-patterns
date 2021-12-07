package train

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/mediator/code/interfaces"
)

type goodsTrain struct {
	mediator interfaces.Mediator
}

func NewGoodsTrain(mediator interfaces.Mediator) *goodsTrain {
	return &goodsTrain{
		mediator: mediator,
	}
}

func (t *goodsTrain) RequestArrival() {
	if t.mediator.CanArrival(t) {
		fmt.Println("GoodsTrain: Arrivals")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (t *goodsTrain) Departure() {
	fmt.Println("GoodsTrain: Departures")
	t.mediator.NotifyFree()
}

func (t *goodsTrain) PermitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Arrivals")
}
