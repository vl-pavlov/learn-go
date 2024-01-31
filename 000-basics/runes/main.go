package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//--	USING RUNES

	// Runes are declared by placing a character inside single quotes
	myRune := '🔥'
	fmt.Printf("My rune: %c; Unicode: %U;", myRune, myRune)

	//--	DECLARATION METHODS

	var letterA rune = 65
	fmt.Printf("Letter %c\n", letterA)

	letterB := rune(66)
	fmt.Printf("Letter %c\n", letterB)

	var letterC int32 = 67
	fmt.Printf("Letter %c\n", letterC)

	letterD := int32(68)
	fmt.Printf("Letter %c\n", letterD)

	var sparkles rune = '\U00002728'
	fmt.Printf("Sparkles %c\n", sparkles)

	thumbsUp := '\U0001F44D'
	fmt.Printf("Thumbs up %c\n", thumbsUp)

	//--	COUNT RUNES

	myString1 := "Hello❗"
	stringLength := len(myString1)
	numberOfRunes := utf8.RuneCountInString(myString1)

	fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)

	//--	TYPE CONVERSION RUNES

	myRuneSlice2 := []rune{'g', 'o', 'l', 'a', 'n', 'g'}
	myString2 := string(myRuneSlice2)
	fmt.Printf("%s\n", myString2)

	myString3 := "Great job! 🔥"
	myRuneSlice3 := []rune(myString3)
	fmt.Printf("%c\n", myRuneSlice3)
}
