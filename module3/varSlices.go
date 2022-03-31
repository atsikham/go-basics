package main

import (
	"fmt"
)

func main() {
	sli1 := make([]int, 10) // init to zero, length == capacity
	fmt.Println(cap(sli1) == len(sli1))
	sli2 := make([]int, 10, 15) // underlying array is bigger than a slice
	fmt.Println(cap(sli2) == len(sli2))

	sli := make([]int, 0, 3)
	fmt.Println(cap(sli))
	sli = append(sli, 100)
	fmt.Println(cap(sli)) // still the same
	fmt.Println(sli)
}
