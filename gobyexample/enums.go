package main

import "fmt"

// enumerated types (enums)
// special case of sum types

// type that has a fixed number of 
// possible values, each with a distinct
// name.

// go doesn't have an enum type but it's
// simple to implement using golang idioms

// our enum:
type ServerState int

const (
	// iota generates successive const values
	// in this case incrementing +1 from 0
	// satisfies ServerState int
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "Retrying",
}

// here we implement the fmt.Stringer() interface
// so enum values are converted to stateName strings
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)	
	
	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}