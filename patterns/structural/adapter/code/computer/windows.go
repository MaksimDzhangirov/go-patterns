package computer

import "fmt"

type windows struct {
}

func NewWindows() *windows {
	return &windows{}
}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}