package main

import (
	"fmt"
	"math"
)

// Interfaces define a named collection of methods
type geometry interface {
	area() float64
	perim() float64
}

// Structs which will implement the geometry interface
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface, with just need to implement
// the methods defined within that interface
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Passing in an interface type means we can call
// that interface's methods against the passed object
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The circle and rect struct types both implement the
	// geometry interface so we can use instances of these
	// structs as arguments to measure.

	measure(r)
	measure(c)
}
