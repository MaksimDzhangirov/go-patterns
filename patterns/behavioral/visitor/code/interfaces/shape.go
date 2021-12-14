package interfaces

type Shape interface {
	GetType() string
	Accept(Visitor)
}