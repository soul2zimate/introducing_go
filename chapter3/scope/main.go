package main

import "fmt"

var x = "Hello, World"

func main() {
	fmt.Println(x)
	f()
}
func f() {
	fmt.Println(x)
}
