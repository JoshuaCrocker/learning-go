package main

import "fmt"

func main() {
	// Channels are unbuffered by default, which means
	// there needs to be a reciever for each message
	// which is sent to te channel. By creating a buffered
	// channel we are able to allow channels to queue up a
	// series of messages, which can be recieved at a later
	// point.
	// This creates a buffered channel which has capacity for
	// two strings.
	messages := make(chan string, 2)

	// We can queue up messages within the channel
	messages <- "buffered"
	messages <- "channel"

	// And recieve them later
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
