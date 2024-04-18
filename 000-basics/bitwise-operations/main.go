package main

import (
	"fmt"
)

func main() {

	//--	BITWISE AND (&)

	resultAnd := 5 & 3

	fmt.Println(resultAnd)
	// Output: 1 (101 & 011 = 001)

	//--	BITWISE OR (|)

	resultOr := 5 | 3

	fmt.Println(resultOr)
	// Output: 7 (101 | 011 = 111)

	//--	BITWISE XOR (^)

	resultXor := 5 ^ 3

	fmt.Println(resultXor)
	// Output: 6 (101 ^ 011 = 110)

	//--	BITWISE SHIFT (<<, >>)

	resultShift := 8 << 1
	fmt.Println(resultShift)
	// Output: 16 (01000 << 1 = 10000)

	resultShift = 8 >> 2     // Right shift by 2 positions
	fmt.Println(resultShift) // Output: 2 (1000 >> 2 = 0010)

	//--	SWAP NUMBERS

	int1 := 10
	int2 := 8

	int1 = int1 ^ int2
	int2 = int1 ^ int2
	int1 = int1 ^ int2

	fmt.Println(int1, int2)

}
