package main

import "fmt"

func main() {
	// exercise 2
	sliceA := make([]int, 3, 9)
	fmt.Println(len(sliceA))

	// exercise 3
	// not including the high index itself
	xx := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(xx[2:5])

	// exercise 4
	listA := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest int
	for index, value := range listA {
		if index == 0 {
			smallest = value
		} else if value < smallest {
			smallest = value
		}
	}
	fmt.Println("smallest is", smallest)
}
