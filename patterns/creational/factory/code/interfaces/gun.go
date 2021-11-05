package interfaces

type Gun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
