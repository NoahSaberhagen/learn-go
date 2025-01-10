// channels are the pipes that connect
// concurrent goroutines

package main

import "fmt"

func main() {
	// channels are typed by values they convey
	messages := make(chan string)

	// send a value using the <- syntax
	go func() { messages <- "ping" }()

	// <-channel receives a value from the channel
	// here we receive "ping" which we sent earlier
	msg := <-messages
	fmt.Println(msg)
}

// by default sends and receives block until both the sender
// and receiver are ready.