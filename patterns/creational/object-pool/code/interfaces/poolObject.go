package interfaces

type PoolObject interface {
	GetID() string // Любой идентификатор, который можно использовать, чтобы отличить двух различных объектов из пула
}
