package main

import "fmt"

func main() {
	// Array.
	// In go, arrays are inmutable.
	var array [4]int
	array[0] = 1
	array[1] = 2

	// len. Number of elemnts of the array.
	// cap. Maximum capacity of element of the array.
	fmt.Println(array, len(array), cap(array))

	// Slices.
	// Slides are mutable.
	// Declaration without quantity
	slice := []int{0, 1, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slices and Arays methods
	fmt.Println(slice[0])   // Print first element of the slide
	fmt.Println(slice[:3])  // Print until 3rd element
	fmt.Println(slice[2:4]) // Print between 2 element and 4
	fmt.Println(slice[4:])  // Print elements after 4

	// Slice append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{4, 5, 6}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
