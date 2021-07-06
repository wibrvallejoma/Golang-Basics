// Golang calls Structs instea of Classes.
package main

import "fmt"

// Struct (Class)
type car struct {
	brand string
	year  int
}

func main() {
	// Way to instance a Struct.
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Another way to instance a Struct.
	var anotherCar car
	anotherCar.brand = "Ferrari"

	fmt.Println(anotherCar) // Ferrari and the Zero value (0)
}
