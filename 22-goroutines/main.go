package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 15; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// To invoke a function within a goroutine, use the 'go'
	// keyword. This runs the function concurrently.
	go f("goroutine")

	// Traditional way of calling functins - synchronous
	f("direct")

	// Anonymous functions can also be called as goroutines
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Wait for the execution to finish (use WaitGroup for a
	// better approach to doing this)
	time.Sleep(time.Second)
	fmt.Println("done")
}
