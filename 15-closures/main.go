package main

import "fmt"

// Go supports anonymous functions, which can be used to
// create closures.
func intSeq() func() int {
	var i int = 0

	// Create and return an anonymous function which
	// encloses the i variable.
	return func() int {
		i++
		return i
	}
}

func main() {
	seq := intSeq()

	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	// The state of 'i' is unique to each function
	seq2 := intSeq()

	fmt.Println(seq2())
	fmt.Println(seq2())
	fmt.Println(seq2())
	fmt.Println(seq2())
}
