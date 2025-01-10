package main

import "fmt"

// go by example:
// zeroval doesn't change the i in main
// but zeroptr does
// because it has a reference to the memory address for
// that variable.


func zeroval(ival int) {
  // in here a new pointer is created for the arg 
  ival = 0
  // in the outside scope, variable that we passed as arg 
  // remains unchanged because it has a different
  // memory address.
}

func zeroptr(iptr *int) {
  // in here, we are changing the value
  // at the memory address of the variable we passed
  // as arg
  *iptr = 0
}

// so we have to explicitly police mutability
// at least when the scope changes.

func main() {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  // here
  // the & syntax give the memory address (pointer)
  // to i
  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}

