package main

import "fmt"

func main() {
	var x [5]int                            // initialized by zero values
	var x1 [5]int = [...]int{1, 2, 3, 4, 5} // array literal

	x[0] = 2

	for i, v := range x {
		fmt.Printf("ind %d, val %d\n", i, v)
	}

	fmt.Println(x1[1])
}
