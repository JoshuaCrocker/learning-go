package main

import "fmt"

func main() {
	// The range method can be used to iterate over the
	// elements within various data structures.

	// Arrays and Slices
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// The index is supplied as the first array value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Maps
	// We get the key and value from the range method this time
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// We can also just get the keys in a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. The
	// first value is the starting byte index of the rune and
	// the second the rune itself.
	// TODO what?
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
