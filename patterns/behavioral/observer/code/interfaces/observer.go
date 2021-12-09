package interfaces

type Observer interface {
	Update(s string)
	GetID() string
}
