package main

import (
	"fmt"
)

func main() {
	// Create string using var
	var part1 = "⌘I love"

	// Create string using shorthand notation
	part2 := "Go"

	// Represent string with `  ` (row string)
	part3 := `Programming`

	// Join parts to make a sentence
	sentence := fmt.Sprintf("%s %s %s", part1, part2, part3)

	// You also can join using concatenation: part1 + part2 + part3

	// Print the sentence
	fmt.Println(sentence)

	// Get number of bytes in the sentence
	// (!) If you want get number of utf-symbols you should use `utf8.RuneCountInString`
	length := len(sentence)

	fmt.Println("Length: ", length)

	// Value of first byte
	// (!) If you want get utf-symbol you should use utf8.DecodeRuneInString
	firstChar := sentence[0]

	fmt.Println("First byte: ", firstChar)

	// Iteration all utf-symbols
	for index, runeValue := range sentence {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	// Great job! Now we know basic string operations
}
