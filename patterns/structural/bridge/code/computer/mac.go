package computer

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/bridge/code/interfaces"
)

type mac struct {
	printer interfaces.Printer
}

func NewMac(printer interfaces.Printer) *mac {
	return &mac{
		printer: printer,
	}
}

func (m *mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *mac) SetPrinter(p interfaces.Printer) {
	m.printer = p
}