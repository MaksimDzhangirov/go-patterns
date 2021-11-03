package interfaces

type Shoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}