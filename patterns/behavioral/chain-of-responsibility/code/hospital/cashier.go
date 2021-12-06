package hospital

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/patient"
)

type cashier struct {
	next interfaces.Department
}

func NewCashier() *cashier {
	return &cashier{}
}

func (c *cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) SetNext(next interfaces.Department) {
	c.next = next
}
