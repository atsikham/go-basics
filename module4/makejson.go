package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a name:")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Enter an address:")
	scanner.Scan()
	address := scanner.Text()

	person := map[string]string{"Name": name, "Address": address}

	barr, _ := json.Marshal(person)

	fmt.Println(barr)
}
