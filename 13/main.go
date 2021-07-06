package main

import (
	"fmt"

	pk "golang_basic_course/13/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hello there")

}
