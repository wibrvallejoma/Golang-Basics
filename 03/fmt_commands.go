package main

import "fmt"

func main() {
	// Printtf
	name := "Frederick"
	last_name := 50
	// %s for String
	// %d for Integer
	// %v for unspecific type
	fmt.Printf("%s is %d years old\n", name, last_name)
	fmt.Printf("%v is %v years old\n", name, last_name)

	// Sprintf. Saves a println type message.
	message := fmt.Sprintf("%v is %v years old", name, last_name)
	fmt.Println(message)

	// Data type.
	// Print data type with %T
	fmt.Printf("last_name: %T ", last_name)
	fmt.Printf("name: %T\n", name)
}
