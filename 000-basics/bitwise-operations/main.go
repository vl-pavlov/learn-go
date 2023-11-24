package main

import "fmt"

func main() {

	//--	Bitwise AND (&)

	resultAnd := 5 & 3

	fmt.Println(resultAnd)
	// Output: 1 (101 & 011 = 001)

	//--	Bitwise OR (|)

	resultOr := 5 | 3

	fmt.Println(resultOr)
	// Output: 7 (101 | 011 = 111)

	//--	Bitwise XOR (^):

	resultXor := 5 ^ 3

	fmt.Println(resultXor)
	// Output: 6 (101 ^ 011 = 110)

	//--	Bitwise Shift Operations (<<, >>):
	resultShift := 8 << 1
	fmt.Println(resultShift)
	// Output: 16 (01000 << 1 = 10000)

	resultShift = 8 >> 2     // Right shift by 2 positions
	fmt.Println(resultShift) // Output: 2 (1000 >> 2 = 0010)

}