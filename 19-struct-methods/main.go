package main

import "fmt"

type rect struct {
	width, height int
}

// Methods can be defined on struct types.
// The area method is said to have a reciever of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can have either a pointer or value type reciever
// It may be advantageous to pass a poiner if you want to
// avoid copying the whole struct for each new method call
// or to allow a method to mutate the given struct.
func (r rect) perim() int {
	return (2 * r.width) + (2 * r.height)
}

func main() {
	// Like with structs, Go automatically resolves passing
	// a pointer or value to the method.
	r := rect{10, 15}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
