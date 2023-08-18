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

//________________________________________________________
// Multiple Return Values

// go has built-in support for multiple return values
// often used in idiomatic Go (i.e. return both result and error values)
func vals() (int, int) {
	return 3, 7
}

//________________________________________________________
// Variadic Functions

// can be called with any number of trailing arguments
// i.e. fmt.Println

// function that will take an arbitrary number of ints as arguments
func sum(nums ...int) {
	// within the function, the type of nums is equivalent to []int
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//________________________________________________________

//________________________________________________________

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//________________________________________________________
	// Multiple Return Values

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	//________________________________________________________
	// Variadic Functions

	// variadic functions can be called in the usual way with individual args
	sum(1, 2)
	sum(1, 2, 3)

	// having a slice with multiple args, we can apply them to a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//________________________________________________________
}
