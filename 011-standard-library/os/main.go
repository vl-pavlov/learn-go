package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//--	HOST NAME

	str, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Host name: %v\n", str)

	//--	ENVIRONMENT VARIABLES

	home := os.Getenv("HOME")
	fmt.Println("Home directory:", home)

	err = os.Setenv("MYVAR", "myvalue")
	if err != nil {
		log.Fatal(err)
	}

	myVar := os.Getenv("MYVAR")
	fmt.Println("MYVAR:", myVar)

	err = os.Unsetenv("MYVAR")
	if err != nil {
		log.Fatal(err)
	}

	myVar = os.Getenv("MYVAR")
	fmt.Println("MYVAR:", myVar)

}
