package printer

import "fmt"

type hp struct {
}

func NewHp() *hp {
	return &hp{}
}

func (p *hp) PrintFile() {
	fmt.Println("Printing by a HP printer")
}
