package main

import "fmt"

// Multiple return values can be defined by wrapping the
// return types in brackets, and separating them with commas.
// This feature is commonly used to return both result and
// error values from functions in Go.
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// We can omit return values we don't want using the
	// blank identifier.
	_, c := vals()
	fmt.Println(c)
}
