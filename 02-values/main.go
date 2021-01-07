package main

import "fmt"

func main() {
	// Strings can be contatenated using the + operator
	fmt.Println("go" + "lang")

	// Arithmetic operations behave as normal
	fmt.Println("1+1=", 1+1)

	// ...with floats too
	fmt.Println("7.0/3.0=", 7.0/3.0)

	// Boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
