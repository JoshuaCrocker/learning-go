package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic Switch
	var i int = 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	}

	// Multiple expressions within the same case statement
	// can be separated by a comma.
	// The default statement defines the case where none of
	// the defined case statements match.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	// A switch with no expression can be used as an if/else
	// statement.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch can be used to determine the type type of
	// an interface value.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
