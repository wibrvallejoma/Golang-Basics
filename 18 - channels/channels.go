package main

import "fmt"

// If you are going to add a channel to a method,
// please specify if you are going to add data to the channel
// or to show/extract the added data.
// <-chan . Extract the data from the channel.
// chan<- Add data to the channel.
func say(text string, c chan<- string) {
	// Add data to the channel
	c <- text
}

func addData(c <-chan string) string {
	// Add data to the channel
	text := <-c
	return text
}

func main() {
	// Make channel with good practices.
	// make(chan dataType, how Many go routines will handle. If you do not specify size, golang will stablish a dinamyc value. Please specify size)
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	// Extract information before the function ends. It waits until it gets the data.
	// With goroutines only it does not waits for the exctraction of the data. It could end before the data extraction.
	fmt.Println(<-c)
}
