package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//--	Validating UTF-8
	// This function can be used to check if a string is valid UTF-8.
	invalidString := string([]byte{0xff, 0xfe, 0xfd})
	validString := utf8.ValidString(invalidString)
	fmt.Printf("Valid string: %t \n", validString)

	//--	Counting Runes in a String
	// This example shows how to count the number of runes in a string, which can be different from its length in bytes due to UTF-8 encoding.

	multibyteString := "Hello, 世界"
	fmt.Printf("Number of bytes in string: %d\n", len(multibyteString))
	fmt.Printf("Number of runes in string: %d\n", utf8.RuneCountInString(multibyteString))

	//--	Decoding a Rune
	// This example demonstrates how to decode the first rune of a string and determine its size in bytes.
	multibyteString = "世界"
	resultingRune, size := utf8.DecodeRuneInString(multibyteString)
	fmt.Printf("First rune: %c, size: %d bytes\n", resultingRune, size)

	//--	Encoding a Rune
	// This example shows how to encode a rune into a UTF-8 byte slice.
	var buf [4]byte
	oneRune := '世'
	byteCount := utf8.EncodeRune(buf[:], oneRune)
	fmt.Printf("Encoded rune: % x, size: %d bytes\n", buf[:byteCount], byteCount)

}
