package main

import (
	"fmt"
)

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {

	var p1 Person     // value
	p2 := new(Person) // pointer

	p1.name = "Joe"
	p2.name = "Jane"

	p3 := Person{name: "Przemek", addr: "Krakow", phone: "+48"}

	fmt.Println(p1.name)
	fmt.Println(*p2)
	fmt.Println(p3)
}
