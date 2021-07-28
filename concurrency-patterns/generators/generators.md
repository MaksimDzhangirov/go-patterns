# Generators

## What is Generator?

A function that returns a channel with data.

```go

package main

import (
	"fmt"
	"time"
)

func main() {
	c := generator()
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("End")
}

func generator() <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Generated value = %d", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	return c
}

```
