package main

import "fmt"

func main() {
	var x = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	var x1 = "hello"
	var y1 = "world"
	fmt.Println(x1 == y1)

	var x2 = "hello"
	var y2 = "hello"
	fmt.Println(x2 == y2)

	x3 := "Hello, World"
	fmt.Println(x3)
	var x4 = "Hello, World"
	fmt.Println(x4)
}
