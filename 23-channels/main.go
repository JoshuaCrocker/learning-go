package main

import "fmt"

func main() {
	// Channels act as pipes which connect concurrent goroutines
	// Messages can be sent into a channel from one routine, and
	// recieved within another.
	// By default, sends and recieves block each thread until both
	// the sender and reciever are ready to perform their
	// respective actions.
	messages := make(chan string)

	// Values can be sent to a channel using the <- operator
	go func() {
		messages <- "ping"
	}()

	// the <-channel syntax recieves a value from a channel.
	msg := <-messages
	fmt.Println(msg)
}
