package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Addr  string
	Phone string
}

func main() {

	p1 := Person{Name: "Przemek", Addr: "Krakow", Phone: "+48"}
	var p2 Person

	barr, err := json.Marshal(p1)

	err = json.Unmarshal(barr, &p2)

	fmt.Println(p2, err)
}
