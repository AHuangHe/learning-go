package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant string"

func main() {
	// Hello World
	fmt.Println("Hello World")

	//________________________________________________________
	// Values / Value types
	fmt.Println("1 + 1.0 = ", 2.0)
	fmt.Println(!true || true && false)

	//________________________________________________________
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

	//________________________________________________________
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

	//________________________________________________________
	// For Loop

	// the only looping construct in Go

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for-loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition (basically like a while loop)
	for {
		fmt.Println("loop unconditionally until break/return")
		break
	}

	// can also continue to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//________________________________________________________
	// If/Else

	// in Go, conditions don't need to be surrounded in parentheses
	// if statements don't need an else statement
	// a statement can precede conditionals, any variables declared are
	// available in the current and all subsequent branches

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// no ternary if --> a ? b : c  (if a is true, then b, else c)
	// need a full if statement even for basic conditions

	//________________________________________________________
	// Switch

	// conditionals across many branches
	// we can use commas to separate multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Monday:
		fmt.Println("It's Monday :(")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternate way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// a type switch compares types instead of values
	// can be used to discover the type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	//________________________________________________________
}
