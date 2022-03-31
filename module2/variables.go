package main

import "fmt"

// variable scopes are defined by blocks
var x1 = 4

// each call f() creates an integer
func f() {
	var y = 5
	fmt.Printf("%d\n", x1)
	fmt.Printf("%d\n", y)
}

func g() {
	fmt.Printf("%d\n", x1)
}

/*
multiline
comment
*/

func main() {
	f()
	g()
}
