package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

// Instead of (a int, b int, c string)
// Set the data type of the same arguments
// at the end (Removing int for a)
// result (a, b int, c string)
func trippleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// func <Name>(args) <dataType>
func returnValue(a int) int {
	return a * a
}

func doubleReturn(a int) (c, d int) {
	return a, a * a
}

func doubleReturnMultiType(a int) (c string, d int) {
	c = "Multiplied of"
	d = a * a
	return c, d
}

func main() {
	fmt.Println("Main func")
	normalFunction("normaFunction")
	trippleArgument(10, 5, "hello")
	fmt.Println(returnValue(5))

	value1, value2 := doubleReturn(50)
	fmt.Println(value1, value2)

	// Must put '_' (underscore) for unused variables
	value3, _ := doubleReturn(500)
	fmt.Println(value3)

	value4, value5 := doubleReturnMultiType(50)
	fmt.Println(value4, value5)
}
