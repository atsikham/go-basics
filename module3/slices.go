package main

import (
	"fmt"
)

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)
	s1 = append(s1, "fff")
	fmt.Println(s1)
	fmt.Println(arr)     // "d" is replaced by the new value -> writing to slice will change an underlying array
	fmt.Println(len(s1)) // length
	fmt.Println(cap(s2)) // capacity
}
