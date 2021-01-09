package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 5)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}

func main() {
	// Channels can be used to synchronise execution across
	// multiple branches of execution. While this works, if
	// we're waiting for multiple goroutines to complete, it
	// may be better to use a WaitGroup.

	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
}
