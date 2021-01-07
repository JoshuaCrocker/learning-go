package main

import "fmt"

func main() {
	// Normal if statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// If statements can exist without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can be declared before the conditionals,
	// which will then be available for all branches of the
	// if-elseif-else block

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Go doesn't have a ternary if (conditon ? true : false)
	// statement, so full if statements are required at all
	// times
}
