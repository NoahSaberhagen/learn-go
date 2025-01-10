package main

import "fmt"

// variadic functions
// any number of trailing arguments
// fmt.Println is a good example of a variadic
// function

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}