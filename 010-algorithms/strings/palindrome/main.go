package main

import "fmt"

func main() {
	word := "abccbaq"
	fmt.Println(isPalindrome(word))
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(word string) bool {
	return word == reverse(word)
}
