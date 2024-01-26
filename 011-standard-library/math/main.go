package main

import (
	"fmt"
	"math"
)

func main() {

	//--	CONSTANTS

	// E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
	// Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
	// Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622

	// Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
	// SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
	// SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
	// SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339

	// Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	// Log2E  = 1 / Ln2
	// Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
	// Log10E = 1 / Ln10

	// Floating-point limit values.
	// Max is the largest finite value representable by the type.
	// SmallestNonzero is the smallest positive, non-zero value representable by the type.
	// MaxFloat32             = 0x1p127 * (1 + (1 - 0x1p-23)) // 3.40282346638528859811704183484516925440e+38
	// SmallestNonzeroFloat32 = 0x1p-126 * 0x1p-23            // 1.401298464324817070923729583289916131280e-45

	// MaxFloat64             = 0x1p1023 * (1 + (1 - 0x1p-52)) // 1.79769313486231570814527423731704356798070e+308
	// SmallestNonzeroFloat64 = 0x1p-1022 * 0x1p-52            // 4.9406564584124654417656879286822137236505980e-324

	// Integer limit values.
	// intSize = 32 << (^uint(0) >> 63) // 32 or 64

	// MaxInt    = 1<<(intSize-1) - 1  // MaxInt32 or MaxInt64 depending on intSize.
	// MinInt    = -1 << (intSize - 1) // MinInt32 or MinInt64 depending on intSize.
	// MaxInt8   = 1<<7 - 1            // 127
	// MinInt8   = -1 << 7             // -128
	// MaxInt16  = 1<<15 - 1           // 32767
	// MinInt16  = -1 << 15            // -32768
	// MaxInt32  = 1<<31 - 1           // 2147483647
	// MinInt32  = -1 << 31            // -2147483648
	// MaxInt64  = 1<<63 - 1           // 9223372036854775807
	// MinInt64  = -1 << 63            // -9223372036854775808
	// MaxUint   = 1<<intSize - 1      // MaxUint32 or MaxUint64 depending on intSize.
	// MaxUint8  = 1<<8 - 1            // 255
	// MaxUint16 = 1<<16 - 1           // 65535
	// MaxUint32 = 1<<32 - 1           // 4294967295
	// MaxUint64 = 1<<64 - 1           // 18446744073709551615

	// Calculate the area of a circle
	radius := 5.0
	area := math.Pi * radius * radius
	fmt.Println("Area of circle:", area)

	//--	Basic arithmetic functions

	// Abs(x float64) float64: Returns the absolute value of a number
	// Max(x, y float64) float64: Returns the larger of two numbers
	// Min(x, y float64) float64: Returns the smaller of two numbers
	// Mod(x, y float64) float64: Returns the remainder of dividing x by y

	// Checking if a number is divisible by another number

	x := 15.0
	y := 5.0

	if math.Mod(x, y) == 0 {
		fmt.Printf("%.2f is divisible by %.2f\n", x, y)
	} else {
		fmt.Printf("%.2f is not divisible by %.2f\n", x, y)
	}

	//--	Trigonometric functions

	// Sin(x float64) float64: Returns the sine of an angle (in radians)
	// Cos(x float64) float64: Returns the cosine of an angle (in radians)
	// Tan(x float64) float64: Returns the tangent of an angle (in radians)
	// Functions for inverse trigonometric functions (asin, acos, atan, etc.) are also available

	// Find the sine of 45 degrees
	angle := 45.0
	radians := angle * (math.Pi / 180)
	sine := math.Sin(radians)
	fmt.Println("Sine of 45 degrees:", sine)

	//--	Exponential and logarithmic functions

	// Exp(x float64) float64: Returns the base-e exponent of x
	// Log(x float64) float64: Returns the natural logarithm (base-e) of x
	// Functions for logarithms with other bases (log10, log2) are available as well

	// Calculating the natural logarithm of a number
	x = 10.0
	fmt.Printf("Natural logarithm of %.2f is %.2f\n", x, math.Log(x))

	//--	Rounding and comparison functions

	// Floor(x float64) float64: Returns the largest integer less than or equal to x
	// Ceil(x float64) float64: Returns the smallest integer greater than or equal to x
	// Round(x float64) float64: Rounds x to the nearest integer.
	// IsNaN(x float64) bool: Checks if x is Not a Number (NaN).

	// Round a number to the nearest integer
	number := 3.14159
	rounded := math.Round(number)
	fmt.Println("Rounded value:", rounded)

	//--	Other functions:
	// Sqrt(x float64) float64: Returns the square root of x.
	// Pow(x, y float64) float64: Returns x raised to the power of y.
	// The math package also provides functions for advanced mathematical operations like hyperbolic functions, Bessel functions, etc.

	// Calculating the square root of a number
	x = 16.0
	fmt.Printf("Square root of %f is %.2f\n", x, math.Sqrt(x))

}
