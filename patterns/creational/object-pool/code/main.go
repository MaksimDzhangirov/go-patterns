package main

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/object-pool/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/object-pool/code/pool"
	"log"
	"strconv"
)

func main() {
	connections := make([]interfaces.PoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &pool.Connection{Id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	objectPool, err := pool.InitPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := objectPool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := objectPool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	err = objectPool.Retrieve(conn1)
	if err != nil {
		log.Fatalf("Pool Retrieve Error: %s", err)
	}
	err = objectPool.Retrieve(conn2)
	if err != nil {
		log.Fatalf("Pool Retrieve Error: %s", err)
	}
}
