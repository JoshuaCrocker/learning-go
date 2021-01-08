package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// This is a basic demonstration of the difference between traditional
// passing by value, and passing a pointer.
// zeroval accepts the ival parameter by value, meaning any changes to
// ival will not be maintained outside the scope of the function.
// zeroptr accepts the iptr parameter as a pointer, which means any
// changes to the value of the pointer will be carried outside of the
// scope of the method
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)
}
