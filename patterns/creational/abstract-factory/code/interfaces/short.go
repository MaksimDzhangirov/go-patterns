package interfaces

type Short interface {
	SetLogo(log string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}