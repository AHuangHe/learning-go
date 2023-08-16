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
	// Arrays

	// a numbered sequence of elements of a specific length
	// less common than slices (below) but more useful in special scenarios
	// by default, an array is zero-valued

	// array that will hold exactly 5 ints
	// the type of elements and length are both part of the array's type
	var arr [5]int
	fmt.Println("emp:", arr)
	// setting value at an index
	arr[4] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])
	fmt.Println("We can print the length with builtin len:", len(arr))

	// declaring and initializing an array in one line
	brr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", brr)

	// multi-dimensional arrays (via types)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//________________________________________________________
	// Slices

	// give a more powerful interface to sequences than arrays
	// unlike arrays, slices are typed only by the elements they contain
	// an uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// to create an empty slice with some length, we use the builtin make
	// by default, the capacity of a new slice is equal to its length
	// if we know it's going to grow, it's possible to pass a capacity explicity
	//   as an additional parameter to make
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// we can set and get values at indexes just like arrays
	s[0] = "a"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0])
	fmt.Println("len:", len(s))

	// we can also append new values to slices with builtin append
	s = append(s, "b")
	s = append(s, "c", "d", "e")
	// there is a gap since s[1] and s[2] were not filled
	fmt.Println("apd:", s, "len:", len(s), "cap:", cap(s))

	// slices can also be copy'd
	csl := make([]string, len(s))
	copy(csl, s)
	fmt.Println("cpy:", csl)

	// we can also slice slices with a range of indexes
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	// we can also declare and init a variable for slice in a single line
	tlc := []string{"g", "h", "i"}
	fmt.Println("dcl:", tlc)

	// slices can be also composed into multiple dimensions
	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD2)

	//________________________________________________________
}
