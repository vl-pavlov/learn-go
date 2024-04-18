package main

import "fmt"

func main() {
	var num int
	var ptr *int

	num = 22
	fmt.Println("Address of num:", &num)
	fmt.Println("Value of num:", num)

	ptr = &num
	fmt.Println("\nAddress of pointer ptr:", ptr)
	fmt.Println("Content of pointer ptr:", *ptr)

	num = 11
	fmt.Println("\nAddress of pointer ptr:", ptr)
	fmt.Println("Content of pointer ptr:", *ptr)

	*ptr = 2
	fmt.Println("\nAddress of num:", &num)
	fmt.Println("Value of num:", num)
}
