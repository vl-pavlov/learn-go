package main

import (
	"fmt"
	"os"
)

func main() {
	// function os.Hostname() example

	str, err := os.Hostname()

	fmt.Println("str: %T, %v\n", str, str)
	fmt.Println("err: %T, %v\n", err, err)

	// str: string, [Name]
	// err: <nil>, <nil>

}
