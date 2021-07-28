# Waiting for goroutines

``` go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	chatter("Bob", &wg)
	chatter("Jane", &wg)
	chatter("Mary", &wg)
	chatter("John", &wg)
	chatter("Make", &wg)
	chatter("Anne", &wg)

	wg.Wait()

	fmt.Println("All chatter have finished")
}

func chatter(s string, wg *sync.WaitGroup) {
	talkTime := rand.Intn(10)+2
	timeUp := time.After(time.Duration(talkTime) * time.Nanosecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timeUp:
				// clean up
				fmt.Println(s, "is quiting after", time.Duration(talkTime))				
				return
			default:
				fmt.Println(s)
				time.Sleep(1 * time.Nanosecond)
			}
		}
	}()
}


```
