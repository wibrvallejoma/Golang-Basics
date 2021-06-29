package main

import "fmt"

func main() {
	// Defer
	// With defer the following line will be executed
	// before the code's block or function dies.
	// It's commonly used to close conextion to a server
	// or close a file. It can be placed at the beggining of
	// the function.
	// Please, Only use 1 defer per function.
	defer fmt.Println("Hello")
	fmt.Println("World")

	// Continue and Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println(i)
			continue
		}

		// Break
		if i == 8 {
			break
		}
	}
}
