package main

import "fmt"

// basic addition function
// requires explicit returns
func plus(a int, b int) int {
	return a + b
}

// when having multiple consecutive parameters of the same type
// you may omit the type name for the like-typed parameters up to the final parameter
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
