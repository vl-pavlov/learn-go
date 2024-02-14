package main

import "fmt"

func main() {
	var perfNum, perfsum int
	perfsum = 0

	perfNum = 6

	for i := 1; i < perfNum; i++ {
		if perfNum%i == 0 {
			perfsum = perfsum + i
		}
	}

	if perfNum == perfsum {
		fmt.Println(perfNum, " is a Perfect Number")
	} else {
		fmt.Println(perfNum, " is Not a Perfect Number")
	}
}
