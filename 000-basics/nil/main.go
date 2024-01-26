package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// There must be sufficient information for
	// compiler to deduce the type of a nil value.
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// These lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// This following line doesn't compile.
	// var _ = nil

	// SIZES OF NIL VALUES

	// The memory layouts of all values of a type are always the same. nil values of the type are not exceptions (assume the zero values of the type can be represented as nil). The size of a nil value is always the same as the sizes of non-nil values whose types are the same as the nil value. But nil values of different kinds of types may have different sizes.

	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16

	//--	GET ELEMENTS FROM NIL MAPS

	fmt.Println((map[string]int)(nil)["key"]) // 0
	fmt.Println((map[int]bool)(nil)[123])     // false
	fmt.Println((map[int]*int64)(nil)[123])   // <nil>

	// RANGE OVER NIL

	for range []int(nil) {
		fmt.Println("Hello")
	}

	for range map[string]string(nil) {
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i)
	}

	// for range chan bool(nil) { // deadlock!
	// 	fmt.Println("Bye")
	// }

}
