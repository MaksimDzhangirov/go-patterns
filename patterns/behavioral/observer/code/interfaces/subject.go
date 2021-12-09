package interfaces

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}
