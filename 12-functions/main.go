package main

import "fmt"

// plus takes two integers and returns the sum
func plus(a int, b int) int {
	// Go requires an explicit return statement
	return a + b
}

// Multiple parameters of the same type only need
// to have the type defined once, at the end of the
// parameters with the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
