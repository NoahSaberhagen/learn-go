package main

import "fmt"

// often used to return both result and err
// values 
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// _ is the blank identifier
	_, c := vals()
	fmt.Println(c)
}