package main

import "fmt"

func main() {
	switch res_module := 50 % 2; res_module {
	case 0:
		fmt.Println("Is Pair")
	default:
		fmt.Println("It's Odd")
	}

	// Without condition
	value := 50
	switch {
	case value > 100:
		fmt.Println("Bigger than 100")
	case value < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("No condition")
	}
}
