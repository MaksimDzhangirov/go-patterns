package interfaces

type Collection interface {
	CreateIterator() Iterator
}
