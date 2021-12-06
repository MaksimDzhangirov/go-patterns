package interfaces

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
