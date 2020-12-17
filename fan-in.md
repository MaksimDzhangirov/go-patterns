# Fan-In

## What in Fan-In?

A mechanism to multiplex several inputs into usually one output. Number of outputs could be more than one but less than number of inputs.

``` go

package main

import (
	"fmt"
	"time"
)

func main() {
	g1 := generator(1)
	g2 := generator(2)
	g3 := generator(3)
	g4 := generator(4)

	result := multiplexer(g1, g2, g3, g4)

	go func() {
		for {
			fmt.Println(<-result)
		}
	}()

	time.Sleep(2 * time.Millisecond)
}

func generator(channelNumber int) chan string {
	out := make(chan string)

	go func() {
		for {
			time.Sleep(400 * time.Microsecond)
			out <- fmt.Sprintf("Generate something for channel %d", channelNumber)
		}
	}()
	return out
}

func multiplexer(in1, in2, in3, in4 <-chan string) chan string {
	out := make(chan string)
	reader := func(in <-chan string) {
		for v := range in {
			out <- v
		}
	}

	go reader(in1)
	go reader(in2)
	go reader(in3)
	go reader(in4)
	return out
}


```
