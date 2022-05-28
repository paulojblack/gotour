package main

import "fmt"

func main() {
	name := "bill"

	// namePointer := &name

	// fmt.Println(namePointer)
	// fmt.Println(&namePointer)
	// printPointer(namePointer)
}

func (s string) overloadString() {
	fmt.Println(s)
}
