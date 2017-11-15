package main

import "fmt"

// exercise 1
func sum(slice []int) int {
	var sum int
	return sum
}

// exercise 2
func halves(x int) (int, bool) {
	return x / 2, x%2 == 0
}

// exercise 3
func variadicFunc(args ...int) int {
	var greatest int
	for i, v := range args {
		if i == 0 || v > greatest {
			greatest = v
		}
	}
	return greatest
}

// exercise 4
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

// exercise 5
func fib(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	return fib(x-1) + fib(x-2)

}

// exercise 11
func swap(x, y *int) {
	// tmp := *x
	// *x = *y
	// *y = tmp
	*x, *y = *y, *x // Play with this magic !
}

func main() {
	fmt.Println("the greatest number is : ", variadicFunc(-1, -2))
	fmt.Println("the fib number of 10 is : ", fib(10))
	x := 1
	y := 2
	fmt.Println("before swap : ", x, y)
	swap(&x, &y)
	fmt.Println("after swap : ", x, y)
}
