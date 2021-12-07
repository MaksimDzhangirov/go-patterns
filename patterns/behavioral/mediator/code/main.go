package main

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/mediator/code/manager"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/mediator/code/train"
)

func main() {
	stationManager := manager.NewStationManager()
	passengerTrain := train.NewPassengerTrain(stationManager)
	goodsTrain := train.NewGoodsTrain(stationManager)
	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}
