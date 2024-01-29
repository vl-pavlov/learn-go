package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//--	SCOPES

	x := 5
	fmt.Println("External variable x:", x)

	if true {
		x := 10
		fmt.Println("Internal variable x:", x)
	}

	fmt.Println("External variable x after internal variable x:", x)

	//

	//--	IF-ELSE

	if h := time.Now().Hour(); h < 12 {
		fmt.Println("Now is AM time.")
	} else if h > 19 {
		fmt.Println("Now is evening time.")
	} else {
		fmt.Println("Now is afternoon time.")
	}

	//-- 	FOR

	// The example will print the integers from 0 to 9.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// It is the same as the while loop in other languages
	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for i < 20 {
		fmt.Println(i)
		i++
	}

	//--	SWITCH-CASE

	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)

		// The if-block, not the fallthrough statement,
		// is the final statement in this branch.
		if true {
			//	 fallthrough // error: not the final statement
		}
	case 5, 6, 7, 8:
		n := 99

		// fallthrough // error: not the final statement
		_ = n
	default:
		fmt.Println(n)
		// fallthrough // error: show up in the final branch
	}

	//--	FOR-RANGE

	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}

	//--	GOTO

	i = 0

Next: // here, a label is declared.
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // execution jumps
	}

	// LABELS

	for i := 90; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Print("The smallest prime number larger than ")
		fmt.Println(i, "is", n)
	}

}

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}
