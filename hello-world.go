package main

import "fmt"

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

	f := "this variable was initialized with the := operator"
	// equivalent to var f string = ""
	// this syntax is only available inside functions
	fmt.Println(f)
}
