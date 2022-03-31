package main

import "fmt"

func main() {
	if true {
		fmt.Printf("Yup\n")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break // or continue
		}
		fmt.Println(i)
	}

	// Infinite loop
	/*for {
		fmt.Println("Hi")
	}*/

	x := 5
	// Break is done automatically
	switch x {
	case 1: // x > 1 (tagless)
		fmt.Printf("case1")
	case 2: // x < -1
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
	}
}
