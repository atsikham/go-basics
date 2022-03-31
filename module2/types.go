package main

import (
	"fmt"
	"strings"
)

func main() {
	// Integer: int and uint
	var x int32 = 1
	var y int16 = 2

	// Convert with T() operation
	x = int32(y)
	fmt.Println(x)

	// Floating point
	var x1 float64 = 123.45
	var y1 float64 = 1.2345e2
	fmt.Println(x1)
	fmt.Println(y1)

	// Complex
	var z complex128 = complex(2, 3)
	fmt.Println(z)

	// Strings
	str := "Hi there"
	// Strings are immutable
	str1 := strings.Replace(str, "Hi", "Hello", 1)
	fmt.Printf("%s\n", str1)

	// iota
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)
	fmt.Println(A, B, C, D, E, F)
}
