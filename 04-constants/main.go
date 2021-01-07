package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean and number types
const s string = "constant"

func main() {
	fmt.Println(s)

	// Constants can be defined anywhere a variable can
	const n = 5000000

	// Constants are defined with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constants don't have a type until they're used,
	// or given a type through an explicit converstion
	fmt.Println(int64(d))

	// The number's type can also be inferred from the context.
	// For example, the math.Sin method takes in a float64 parameter
	fmt.Println(math.Sin(d))
}
