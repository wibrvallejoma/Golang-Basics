package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// To make an specific function to a struct
// Must instance the struct before the func name.
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	// & Access the memory location.
	// * Access to the memory address values
	a := 50
	b := &a // Pointer

	fmt.Println(a)  // 50
	fmt.Println(b)  // 0xc0000160a8
	fmt.Println(*b) // 50

	*b = 100

	fmt.Println(a)  // 50
	fmt.Println(*b) // 50

	myPc := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)

	myPc.duplicateRAM()
	fmt.Println(myPc)

	myPc.duplicateRAM()
	fmt.Println(myPc)
}
