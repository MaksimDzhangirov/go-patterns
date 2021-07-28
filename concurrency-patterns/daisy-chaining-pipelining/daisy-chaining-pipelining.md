# Daisy-Chaining or Pipelining

## What is 'Daisy-Chaining' or 'Pipelining'?

A method of linking goroutines that consumes and produces. Can be thought of as an assembly line.

``` go
package main

import (
	"fmt"
)

const N = 20000

func main() {
	c := make(chan int)
	in := c
	var out chan int

	for i := 0; i < N; i++ {
		out = worker(in)
		in = out
	}

	c <- 0
	fmt.Println(<-out)
}

func worker(in chan int) chan int {
	out := make(chan int)

	go func() {
		v := <-in
		out <- v + 1
	}()

	return out
}

```
