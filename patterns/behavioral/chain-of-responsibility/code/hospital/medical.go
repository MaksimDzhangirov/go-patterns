package hospital

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/patient"
)

type medical struct {
	next interfaces.Department
}

func NewMedical() *medical {
	return &medical{}
}

func (m *medical) Execute(p *patient.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *medical) SetNext(next interfaces.Department) {
	m.next = next
}