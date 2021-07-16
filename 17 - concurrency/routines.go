package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// WaitGroup: Manipulates go routines as primitive method (More efficient but hardest to mantain)
	// Acumulates go routines and it release them slowly
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	// Anonymous function
	go func(text string) {
		fmt.Println(text)
	}("Bye")

	time.Sleep(time.Second * 1)
}
