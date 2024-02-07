package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	myList := list.New()

	elm4 := myList.PushBack(4)
	// mylist = 1

	elm1 := myList.PushFront(1)
	// mylist = 1, 4

	myList.InsertBefore(3, elm4)
	// mylist = 1, 3, 4

	myList.InsertAfter(2, elm1)
	// mylist = 1, 2, 3, 4

	myList.Remove(myList.Front())
	// mylist = 2, 3, 4

	myList.Remove(myList.Back())
	// mylist = 2, 3

	// Iterate through list and print its contents.
	for elm := myList.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value)
	}

	// You can also implement Stack and Queue using this package
}
