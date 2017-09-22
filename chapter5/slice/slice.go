package main

import "fmt"

func main() {
	x := make([]float64, 5)
	for _, value := range x {
		fmt.Println(value)
	}
	fmt.Println("=====")

	x1 := make([]float64, 5, 10)
	for _, value := range x1 {
		fmt.Println(value)
	}
	fmt.Println("=====")

	arr := [5]float64{1, 2, 3, 4, 5}
	x2 := arr[0:5]
	for _, value := range x2 {
		fmt.Println(value)
	}
	fmt.Println("=====")

	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)
}
