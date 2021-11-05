package interfaces

type Node interface {
	Print(string)
	Clone() Node
}
