package interfaces

type Train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}
