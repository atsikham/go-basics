package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dat := "Hello, world"
	// The whole file processing, no need to work with descriptors
	err := ioutil.WriteFile("out.txt", []byte(dat), 0777)
	fmt.Println(err)

	f, err := os.Open("out.txt")
	barr := make([]byte, 10)
	_, err = f.Read(barr)
	f.Close()
	fmt.Println(string(barr)) // will print only the text limited by array size

	f, err = os.Create("outfile.txt")
	barr = []byte{1, 2, 3}
	_, err = f.Write(barr)
	_, err = f.WriteString("Hi")
}
