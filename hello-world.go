package main

import (
	"fmt"
	"math"
)

const s string = "constant string"

func main() {
	// Hello World
	fmt.Println("Hello World")

	// Values / Value types
	fmt.Println("1 + 1.0 = ", 2.0)
	fmt.Println(!true || true && false)

	// Variables
	var a = "this is a var"
	fmt.Println(a)

	var b, c int = 1, 2 // int is not necessary (go does type inference)
	fmt.Println(b, c)

	var d int // zero-value initialization (0 in this case since its an int)
	fmt.Println(d)

	e := "this variable was initialized with the := operator"
	// equivalent to var f string = ""
	// this syntax is only available inside functions
	fmt.Println(e)

	// Constants
	fmt.Println(s)
	// constant expressions perform arithmetic with arbitrary precision
	const n = 500000000

	// a numeric constant has not type until it's given one
	const f = 3e20 / n
	// no type
	fmt.Println(f)
	// given type
	fmt.Println(int64(f))

	fmt.Println(math.Sin(n))

}
