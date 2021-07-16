package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// Override Print String
// Structs have a String method
// wich one can be overrided to show
// a different output.
func (myPC pc) String() string {
	return fmt.Sprintf("I have %d GB RAM, %d GB Disk and is a %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPc)
}
