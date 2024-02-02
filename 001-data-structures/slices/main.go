package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0
	var slice []string
	fmt.Println("Uninitialized:", slice, slice == nil, len(slice) == 0)

	// To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 5 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	slice = make([]string, 5)
	fmt.Println("Empty slice:", slice, "Len:", len(slice), "Cap:", cap(slice))

	// We can set and get just like with arrays.
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("Our slice:", slice)
	fmt.Println("Get [2]:", slice[2])

	// `len`` returns the length of the slice as expected.
	fmt.Println("Len:", len(slice))

	// In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin `append`, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("Appended:", slice)

	// Slices can also be copied. Here we create an empty slice c of the same length as s and copy into `copied` from `slice`.
	copied := make([]string, len(slice))
	copy(copied, slice)
	fmt.Println("Copied:", copied)

	// Slices support a `slice` operator with the syntax slice[low:high]. For example, this gets a slice of the elements [2], [3], and [4].
	sliced1 := slice[2:5]
	fmt.Println("Slice [2:5]:", sliced1)

	// This slices up to (but excluding) [5].
	sliced2 := slice[:5]
	fmt.Println("Slice [:5]:", sliced2)

	// And this slices up from (and including) [2].
	sliced3 := slice[2:]
	fmt.Println("Slice [2:]:", sliced3)

	// We can declare and initialize a variable for slice in a single line as well.
	letters1 := []string{"g", "h", "i"}
	fmt.Println("Letters 1:", letters1)

	// The slices package contains a number of useful utility functions for slices.
	letters2 := []string{"g", "h", "i"}
	fmt.Println("Letters 2:", letters1)
	if slices.Equal(letters1, letters2) {
		fmt.Println("Letters 1 == Letters 2")
	}

	//--	2D SLICE

	// Slices can be composed into multi-dimensional data structures.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d slice: ", twoD)

	//--	SLICE TO ARRAY

	// Convert slice to array
	array := [2]string{}

	copy(array[:], slice)
	fmt.Println("Array: ", array, " from ", slice)

}
