package main

var x int = 1 // type can be avoided here
var y int
var ip *int

func main() {
	ip = &x         // get pointer
	y = *ip         // set value by pointer
	ptr := new(int) // create a pointer
	*ptr = 3        // set value by pointer
}
