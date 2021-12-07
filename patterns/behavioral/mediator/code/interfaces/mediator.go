package interfaces

type Mediator interface {
	CanArrival(train Train) bool
	NotifyFree()
}
