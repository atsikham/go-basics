package main

import (
	"fmt"
)

func main() {
	var idMap map[string]int
	// Create an empty map
	idMap = make(map[string]int)

	// Or define a map literal in place
	idMap1 := map[string]int{"Joe": 123}

	fmt.Println(idMap)

	idMap["Jane"] = 345

	fmt.Println(idMap)

	delete(idMap, "Jane")
	delete(idMap, "Joe") // no error if key doesn't exist

	fmt.Println(idMap)

	id, present := idMap1["Joe"]
	fmt.Println(id, present)

	id1, present1 := idMap1["Jane"]
	fmt.Println(id1, present1)

	fmt.Println(idMap1["Joe"])

	for key, val := range idMap1 {
		fmt.Println(key, val)
	}
}
