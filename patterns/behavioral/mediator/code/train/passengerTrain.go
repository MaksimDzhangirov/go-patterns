package train

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/mediator/code/interfaces"
)

type passengerTrain struct {
	mediator interfaces.Mediator
}

func NewPassengerTrain(mediator interfaces.Mediator) *passengerTrain {
	return &passengerTrain{
		mediator: mediator,
	}
}

func (t *passengerTrain) RequestArrival() {
	if t.mediator.CanArrival(t) {
		fmt.Println("PassengerTrain: Arrivals")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (t *passengerTrain) Departure() {
	fmt.Println("PassengerTrain: Departures")
	t.mediator.NotifyFree()
}

func (t *passengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Arrivals")
}
