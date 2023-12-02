package main

import (
	"fmt"
	"strings"
)

const testString = "Mary had a little lamb"

func main() {

	// Checking for a substring
	substr := "lamb"
	contain := strings.Contains(testString, substr)
	fmt.Printf("The '%s' contains '%s': %t \n", testString, substr, contain)

	substr = "wolf"
	contain = strings.Contains(testString, substr)
	fmt.Printf("The '%s' contains '%s': %t \n", testString, substr, contain)

	// Checking for prefixes and suffixes
	prefix := "Mary"
	starts := strings.HasPrefix(testString, prefix)
	fmt.Printf("The '%s' starts with '%s': %t \n", testString, prefix, starts)

	suffix := "lamb"
	ends := strings.HasSuffix(testString, suffix)
	fmt.Printf("The '%s' ends with '%s': %t \n", testString, suffix, ends)

	// Counting instances of a substring
	substr = "e"
	strings.Count("cheese", substr)

	// Comparing strings
	strings.EqualFold("Go", "go")

	// Replacing substrings
	strings.Replace("oink oink oink", "k", "ky", 2)

	// Spliting a string into a slice of substrings based on a separator
	phrase := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(phrase, " ")
	fmt.Println("Words:", words)

	// Joining the words back into a string, separated by a comma
	commaSeparated := strings.Join(words, ", ")
	fmt.Println("Comma-separated:", commaSeparated)

	// Converting the phrase to uppercase
	upper := strings.ToUpper(phrase)
	fmt.Println("Uppercase:", upper)

	// Removing leading and trailing white space from a string
	strWithSpaces := "   Go Programming   "
	trimmed := strings.TrimSpace(strWithSpaces)
	fmt.Println("Trimmed:", trimmed)

}
