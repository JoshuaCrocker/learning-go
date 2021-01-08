package main

import "fmt"

// Structs group together multiple variables within a single
// type. For example, the below person struct has a string name
// and an integer age.
type person struct {
	name string
	age  int
}

// This function constructs a new person struct with the given
// name. Returning the reference struct is okay, as it will
// continue to exist outside the scope of the function.
func newPerson(name string) *person {
	var p person = person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Create a new struct
	fmt.Println(person{"John", 20})

	// Fields are populated by naming them
	fmt.Println(person{name: "Joe", age: 20})

	// Fields which are omitted are zero-values
	fmt.Println(person{name: "Zach"})

	// An ampersand prefix returns a pointer to the struct
	fmt.Println(&person{name: "Ann"})

	// It's traditional to wrap the creation of new structs
	// within a constructor function
	fmt.Println(newPerson("Jon"))

	// Fields can be accessed with a full-stop
	s := person{name: "Sean"}
	fmt.Println(s.name)

	// Full-stops can also be used to reference fields within
	// struct pointers
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
