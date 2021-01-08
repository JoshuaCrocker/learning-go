package main

import (
	"errors"
	"fmt"
)

// Errors in go are traditionally returned using a separate
// return value. This is contrasting to other languages,
// which use exceptions to indicate errors.

// Errors are traditionally the final return value, and have the
// type 'error'
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a new error with the given error
		// string
		return -1, errors.New("can't work with 42")
	}

	// A nil value is used to indicate there was no error
	return arg + 3, nil
}

// It is possible to create custom errors by implementing the Error()
// method on them
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// In this case we use &argError syntax to build a new struct,
		// supplying values for the two fields arg and prob.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// This loop shows the error returning at work.
	for _, i := range []int{7, 42} {
		// Checking for errors within the if statement is a commonly
		// used technique within Go
		//                VVVVVVVV
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
}
