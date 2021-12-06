package hospital

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/patient"
)

type reception struct {
	next interfaces.Department
}

func NewReception() *reception {
	return &reception{}
}

func (r *reception) Execute(p *patient.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *reception) SetNext(next interfaces.Department) {
	r.next = next
}