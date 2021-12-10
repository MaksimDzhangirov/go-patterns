package interfaces

type Cache interface {
	SetEvictionAlgo(e EvictionAlgo)
	Add(key, value string)
	Get(key string)
	Evict()
}

type EvictionAlgo interface {
	Evict(c Cache)
}
