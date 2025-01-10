// you can specify if a channel is meant to only send or 
// receive values

package main

import "fmt"

// put msg in pings channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// get msg from pings
// send to pongs channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// create channels
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}