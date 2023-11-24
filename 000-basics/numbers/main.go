package main

import "fmt"

func main() {

	//--	Integers

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// byte        alias for uint8
	// rune        alias for int32

	// 42
	// 4_2
	// 0600
	// 0_600
	// 0o600
	// 0o600 // second character is capital letter 'O'
	// 0xBadFace
	// 0xBad_Face
	// 0x_67_7a_2f_cc_40_c6
	// 170141183460469231731687303715884105727
	// 170_141183_460469_231731_687303_715884_105727

	// strconv.Atoi() function converts a string to an integer.

	// strconv.ParseInt() can parse numbers in different bases, including hexadecimal.

	// strconv.ParseUint() can parse unsigned integers in different bases.

	//--	Floating-Point Numbers

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// https://www.h-schmidt.net/FloatConverter/IEEE754.html

	//
	// 0.
	// 72.40
	// 072.40 // == 72.40
	// 2.71828
	// 1.e+0
	// 6.67428e-11
	// 1e6
	// .25
	// .12345e+5
	// 1_5.      // == 15.0
	// 0.15e+0_2 // == 15.0

	// 0x1p-2      // == 0.25
	// 0x2.p10     // == 2048.0
	// 0x1.Fp+0    // == 1.9375
	// 0x.8p-0     // == 0.5
	// 0x_1FFFp-16 // == 0.1249847412109375
	// 0x15e - 2   // == 0x15e - 2 (integer subtraction)

	//The strconv.ParseFloat() function can be used to convert a string to a floating-point number.

	//--	Complex numbers
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	z := 3 + 4i
	fmt.Println("Real part:", real(z))
	fmt.Println("Imaginary part:", imag(z))

	//
	var realPart float64 = 3.0
	var imagPart float64 = 4.0
	c := complex(realPart, imagPart)
	fmt.Println("Complex number:", c)

	//--	Arithmetic Operations in Go

	a := 3.5
	b := 0.5

	// Addition
	sum := a + b
	fmt.Println("Sum: ", sum)

	// Subtraction
	difference := a - b
	fmt.Println("Difference: ", difference)

	// Multiplication
	product := a * b
	fmt.Println("Product: ", product)

	// Division
	quotient := a / b
	fmt.Println("Quotient: ", quotient)

	// Go also provides compound assignment operators like +=, -=, *=, /= for shorthand arithmetic operations combined with assignment.

	// Using remainder operator %
	fmt.Println(10 % 3) // Output: 1 (10 divided by 3 leaves a remainder of 1)
	fmt.Println(15 % 4) // Output: 3 (15 divided by 4 leaves a remainder of 3)

}
