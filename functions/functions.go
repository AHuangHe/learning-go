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
// Closures and Anonymous Functions

// Go supports anonymous functions, which can form closures
// useful for when you want an inline function without having to name it

// this function returns another function, which we define anonymously in the body
// the returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

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
	// Closures and Anonymous Functions

	// we call intSeq, assigning the result to nextInt
	nextInt := intSeq()
	// this function value captures its own i value,
	// which will be updated each time we call nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique to that particular function
	// the following function will have its own value and state
	newInts := intSeq()
	fmt.Println(newInts())
	//________________________________________________________
}
