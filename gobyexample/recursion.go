package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	fmt.Println(fact(7))

	// anonymous functions can
	// also be recursive but it requires
	// explicitly storing the function
	// by assigning it with var	
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n 
		}
		return fib(n - 1) + fib(n - 2)
	}
	
	fmt.Println(fib(7))
}