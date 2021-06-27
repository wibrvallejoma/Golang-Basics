package main

import "fmt"

// Verify if a number is pair.
func isPair(n int) bool {
	if n%2 == 0 || n == 1 {
		return true
	} else {
		return false
	}
}

// Verify if two strings are the same to others two.
func stringsAreEquals(a, b string) bool {
	c := "Hello"
	d := "World"
	if a == c && b == d {
		return true
	} else {
		return false
	}
}

func main() {
	// Challenges
	// Verify if a number is par
	// Verify if two strings are the same to other two
	a := 3
	b := 10
	fmt.Printf("Is Pair (%d, %d)\n", a, b)
	fmt.Printf("(%d): (%t)\n", a, isPair(a))
	fmt.Printf("(%d): (%t)\n", b, isPair(b))

	// Equals strings
	c := "Hell"
	d := "wo"
	fmt.Printf("Equals strings (%s, %s)\n", c, d)
	fmt.Printf("(%s, %s): (%t)", c, d, stringsAreEquals(c, d))
}
