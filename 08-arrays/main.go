package main

import "fmt"

func main() {
	// Arrays can be initialised with a type, and the number
	// of items that array holds. Arrays created with no values
	// are zero-values, which is 0 for integers.
	var a [5]int
	fmt.Println(a)

	// Setting and getting is like other languages
	// Arrays are zero-indexed
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len() returns the length of an array
	fmt.Println("len:", len(a))

	// Arrays can be declared and initialised in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Two-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
