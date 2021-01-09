package main

import "fmt"

func main() {
	// Basic channel operations are blocking by default.
	// We can use a switch statement with a default clause
	// to stop channel operations from blocking the current
	// thread.

	messages := make(chan string)
	signals := make(chan bool)

	// This is a non-blocking recieve statement. If there is
	// a message then the first case will be run, otherwise
	// the default case will be run.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// This is a non-blocking send statement. This will fail
	// and return the default branch, as there is no buffer
	// on the channel, and no reciever is set up to recieve the
	// message.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple case blocks to set up a multi-way
	// non-blocking series of channel operations.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
