package main

import "fmt"

// The ping function on accepts a channel for sending values.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and
//  a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// When using channels as function parameters, the direction
	// of the channel can be specified (i.e. can we only send or
	// receive with that channel reference). This improves the
	// type safety of the system.

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
