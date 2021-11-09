package computer

import "fmt"

type mac struct {
}

func NewMac() *mac {
	return &mac{}
}

func (m *mac) InsertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}
