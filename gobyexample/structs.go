package main

import "fmt"

type person struct {
  name string
  age int
}

// & - address of operator
//    returns memory address of variable
// * - dereference operator
//    accesses value stored at memory
//    address pointed to
func newPerson(name string) *person {
  p := person{name: name}
  p.age = 42

  // the way pointers are being used
  // each time it runs a new person
  // is created...?

  return &p
}

func main() {
  fmt.Println(person{"bob", 20})

  fmt.Println(person{name: "alice", age: 30})

  fmt.Println(person{name: "fred"})

  fmt.Println(&person{name: "ann", age: 40})

  fmt.Println(newPerson("jon"))

  s := person{name: "sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)

  dog := struct {
    name string
    isGood bool
  }{
    "Rex",
    true,
  }

  fmt.Println(dog)
}
