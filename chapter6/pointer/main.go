package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPrt *int) {
	*xPrt = 1
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
