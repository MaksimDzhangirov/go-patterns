# Graceful Goroutine Exit

## Gracefully terminating a 'goroutine'
Signal a goroutine to cleanup and exit. Using input channel, quit channel to send signal(s) to the goroutine.

``` go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c0 := chatter("Bob", quit)
	c1 := chatter("John", quit)
	c2 := chatter("Jane", quit)
	c3 := chatter("Mary", quit)

	go func() {
		for {
			select {
			case v := <-c0:
				fmt.Println(v)
			case v := <-c1:
				fmt.Println(v)
			case v := <-c2:
				fmt.Println(v)
			case v := <-c3:
				fmt.Println(v)
			}
		}

	}()

	time.Sleep(30 * time.Nanosecond)
	quit <- true
	quit <- true
	quit <- true
	quit <- true
}

func chatter(s string, quit <-chan bool) <-chan string {
	out := make(chan string)

	go func() {
		fmt.Printf("%s started\n", s)
		loop := true
		for loop {
			select {
			case <-quit:
				// do some clean up work
				fmt.Printf("%s is quiting\n", s)
				loop = false
				break
			case out <- s:
				time.Sleep(time.Duration(rand.Intn(5)) * time.Nanosecond)
			}

		}

	}()

	return out
}

```
