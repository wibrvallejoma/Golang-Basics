package main

import "fmt"

func main() {
	// map[<key data type>]<values data type>

	// Maps are concurrent. So it could show
	// the values more random than usual.
	// Take care of it.
	m := make(map[string]int)

	m["Jose"] = 14
	m["Manuel"] = 20

	fmt.Println(m)

	// Range map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Obtain certain value
	value := m["Jose"]
	fmt.Println(value)

	// If the key doesn't exits.
	// Golang will show the Zero value
	// for that key instead.
	// In this case the zero value of int is 0.
	value2 := m["Joseph"]
	fmt.Println(value2)

	// Get Value and boolean if that Key exists.
	j_value, ok := m["Joshep"]
	fmt.Println(j_value, ok) // False because doesn't exists in dict.
}
