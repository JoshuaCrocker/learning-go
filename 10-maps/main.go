package main

import "fmt"

func main() {
	// Maps are the associative data type in Go
	// Create an empty map
	m := make(map[string]int)

	// Maps can be written to using the normal, array syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Values can be retrieved using the normal, array syntax
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len() returns the number of key-value pairs in the map
	fmt.Println("len:", len(m))

	// delete() removes the specified pair from the map
	delete(m, "k2")
	fmt.Println("map:", m)

	// The second return value indicates the presence of the
	// given key within the map. This is useful to differentiate
	// between zero values (map values initiated to 0 or "") and
	// keys which don't exist.
	// As we don't want the value from the map, we use the blank
	// identifier (_) to ignore it.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// One-line map initialisation
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
