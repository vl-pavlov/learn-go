package main

import "fmt"

func main() {
	var factorsnum, i int

	factorsnum = 10

	fmt.Println("The Factors of the ", factorsnum, " are = ")
	for i = 1; i <= factorsnum; i++ {
		if factorsnum%i == 0 {
			fmt.Println(i)
		}
	}
}
