// by default channels are unbuffered

// meaning they will only accept sends if there is a corresponding
// receive ready to receive the sent value


// buffered channels accept a limited number of values without a 
// corresponding receiver for those


package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"	

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// so a buffered channel is a fifo list