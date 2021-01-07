package main

import "fmt"

// The for loop is the only type of loop within the Go language

func main() {
	// Basic For Loop
	var i = 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// "Traditional" for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// An infinite loop can be made by creating a for loop
	// without any conditions.
	// Infinite loops can end using the 'break' or 'return' keyword.
	var k = 0
	for {
		fmt.Println("loop")

		if k > 10 {
			break
		}

		k++
	}

	// The continue keyword can be used to go to the next iteration
	// of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
