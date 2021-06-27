package main

import "fmt"

func main() {
	// Square Area
	const squareBase = 10
	SquareArea := squareBase * squareBase
	fmt.Println("Square area:", SquareArea)

	x := 10
	y := 50

	// Addition
	result := x + y
	fmt.Println("Addition:", result)

	// Substraction
	result = x - y
	fmt.Println("Substraction:", result)

	// Multiplcation
	result = x * y
	fmt.Println("Multiplcation:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Mod
	result = y % x
	fmt.Println("Mod:", result)

	// Increment
	x++
	fmt.Println("Increment:", x)

	// Decrement
	y--
	fmt.Println("Decrement:", y)

	// Challenge Areas:
	// -Rect, Trapezium, Circle

	// Rect Area
	const height float64 = 20
	const width float64 = 40
	const rect_area float64 = height * width
	fmt.Println("Rect")
	fmt.Println("Sizes:", height, width)
	fmt.Println("Area:", rect_area)

	// Trapezium Area
	const a float64 = 20
	const b float64 = 40
	const height_t float64 = 30
	const trapez_area float64 = (a + b) / 2 * height_t
	fmt.Println("Trapezium")
	fmt.Println("Sizes:", a, b, height_t)
	fmt.Println("Area:", trapez_area)

	// Circle Area
	const PI float64 = 3.1416
	const r float64 = 5
	const circle_area float64 = PI * (r * r)
	fmt.Println("Circle")
	fmt.Println("Size:", r)
	fmt.Println("Area:", circle_area)

}
