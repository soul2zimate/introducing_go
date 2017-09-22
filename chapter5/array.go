package main

import "fmt"

func main() {
	array()
	average()
	rangeFunc()
}

func array() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println("last element is", x[len(x)-1])
}

func average() {
	x := [5]float64{98, 93, 77, 82, 83}
	total := 0.0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func rangeFunc() {
	x := [5]float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
