package pool

import (
	"errors"
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/object-pool/code/interfaces"
	"sync"
)

type pool struct {
	idle     []interfaces.PoolObject
	active   []interfaces.PoolObject
	capacity int
	mulock   *sync.Mutex
}

// InitPool инициализирует пул
func InitPool(poolObjects []interfaces.PoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, errors.New("cannot create a pool of 0 length")
	}
	active := make([]interfaces.PoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

// Loan заимствуем объект из пула
func (p *pool) Loan() (interfaces.PoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, errors.New("no pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.GetID())
	return obj, nil
}

// Retrieve возвращаем объект в пул
func (p *pool) Retrieve(target interfaces.PoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.Remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.GetID())
	return nil
}

// Remove удаляем объект из пула
func (p *pool) Remove(target interfaces.PoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.GetID() == target.GetID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return errors.New("target pool object doesn't belong to the pool")
}
