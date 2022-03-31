package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var persons []Person

	scannerStdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a file name:")
	scannerStdin.Scan()
	fileName := scannerStdin.Text()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scannerFile := bufio.NewScanner(file)
	for scannerFile.Scan() {
		personInfo := strings.Split(scannerFile.Text(), " ")
		person := Person{personInfo[0], personInfo[1]}
		persons = append(persons, person)
	}

	for _, person := range persons {
		fmt.Printf("First name: %s, Last name: %s\n", person.fname, person.lname)
	}
}
