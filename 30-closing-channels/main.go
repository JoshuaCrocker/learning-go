package main

import "fmt"

func main() {
	// Closing a channel indicates to the recievers that there will
	// no longer be any messages sent on that channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// This goroutine loops until the jobs channel has been closed
	// and all values have already been read. It does this using the
	// special two-value return when recieving messages from a
	// channel. In this instance, the first value is the contents of
	// the message, and the second value is a boolean flag which
	// will be false if all messages have been read AND the channel
	// has been closed.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
