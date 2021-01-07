package main

import "fmt"

func main() {
	// Slices are more powerful than arrays as they're fixed to
	// a specific type, but they can be variable length. The make()
	// function can be used to create a new slice, with the second
	// parameter being the initial length of the slice (all zero valued)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Array-style getting and setting
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// len() returns the length of the slice
	fmt.Println("len:", len(s))

	// There are various methods which can act upon the slice, such
	// as the append method, which takes the slice and the new items,
	// and returns a new slice containing the new elements.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copied.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices can be 'sliced' up to only get some of the elements. This is
	// done with the array-get syntax, but with [low:high(non-inclusive)]
	// instead of just one number.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Slice up to, but excluding element s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// Slices from s[2] to the end of the slice
	l = s[2:]
	fmt.Println("sl3:", l)

	// Slices can be declared and initialised in one line, much like arrays
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be made into multi-dimensional structures, where each of the
	// inner slices, unlike arrays, can be different lengths.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
