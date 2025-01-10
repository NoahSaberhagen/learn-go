package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// everthing with go is being run concurrently
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// `go run goroutines.go`
//
// direct : 0
// direct : 1
// direct : 2
//
// --- output at this point may change as the goroutines are concurrent
//
// goroutine : 0
// going
// goroutine : 1
// goroutine : 2
// done