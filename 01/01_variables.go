package main

import (
	"fmt"
)

func main() {
	// Constants declaration Option 1.
	const pi float64 = 3.14
	// Constants declaration Option 2.
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Variables declaration.
	base := 12 // If variable doesn't have ':' that means that the variable already exists.
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values. Instead of var = Null, GoLang does it automatically.
	var a int     // 0
	var b float64 // 0
	var c string  // ''
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Square Area
	const squareBase = 10
	SquareArea := squareBase * squareBase
	println("Square area:", SquareArea)
}
