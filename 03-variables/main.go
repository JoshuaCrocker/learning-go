package main

import "fmt"

func main() {
	// Variable Declaration is done with the 'var' keyword
	var a = "string"
	fmt.Println(a)

	// Multiple variables can be declared at once
	var b, c = 1, 2
	fmt.Println(b, c)

	// Go infers the type of variables
	var d = true
	fmt.Println(d)

	// The type of variables can be explicitly set
	var e int = 1
	fmt.Println(e)

	// Variables can be declared and set using the := shorthand
	// ...is the same as var f string = "value"
	f := "value"
	fmt.Println(f)
}
