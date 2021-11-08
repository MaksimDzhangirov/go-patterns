package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

/*
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
*/

func getInstance() *single {
	var once sync.Once
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating Single Instance Now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
