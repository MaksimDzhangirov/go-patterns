package computer

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/bridge/code/interfaces"
)

type windows struct {
	printer interfaces.Printer
}

func NewWindows(printer interfaces.Printer) *windows {
	return &windows{
		printer: printer,
	}
}

func (w *windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *windows) SetPrinter(p interfaces.Printer) {
	w.printer = p
}
