package printer

import "fmt"

type epson struct {
}

func NewEpson() *epson {
	return &epson{}
}

func (p *epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
