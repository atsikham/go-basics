package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	inputLower := strings.ToLower(line)
	if strings.HasPrefix(inputLower, "i") && strings.HasSuffix(inputLower, "n") && strings.Contains(inputLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
