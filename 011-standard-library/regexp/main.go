package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "We 🔥Love🔥 Go **Programming** **Language^^"

	// Regular expression to extract all Non-Alphanumeric Characters from a String
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern: %v\n", re.String())
	fmt.Println(re.MatchString(str1))

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
