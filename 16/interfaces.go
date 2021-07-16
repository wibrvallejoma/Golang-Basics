// Interfaces can create a method for multiples Structs.
// Instead of creating the same method for each Struct.
package main

import "fmt"

type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figure2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// List of interfaces
	myInterfaces := []interface{}{"Hi", 12, 4.90} // {} First brackets = Blank. Second brackets = Initial data
	fmt.Println(myInterfaces...)
}
