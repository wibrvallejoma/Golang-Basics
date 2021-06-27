package main

import "fmt"

func main() {

	// Conditional For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	// Basic for each approach
	numbersList := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, n := range numbersList {
		fmt.Printf("index %d, number: %d \n", i, n)
	}
}
