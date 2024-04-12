package main

import "fmt"

func main() {
	list := NewList()

	list.Add(NewNumber(1))
	list.Add(NewNumber(1.2))

	// Error
	// list.Add(NewNumber("1.2"))

	fmt.Println("Sum: ", list.Sum())
}
