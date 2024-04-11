package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// create a byte array of a string
	data := []byte("Hello Gopher!")

	// write data to a hello file, with 0777 file permission
	err := os.WriteFile("hello.txt", data, 0777)

	if err != nil {
		log.Fatalf("%v", err)
	}

	// read the hello.txt
	content, err := os.ReadFile("hello.txt")

	if err != nil {
		log.Fatalf("error while reading %v", err)
	}

	// convert the byte into string
	fmt.Println(string(content))
}
