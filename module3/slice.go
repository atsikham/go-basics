package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	scanner := bufio.NewScanner(os.Stdin)
Loop:
	for {
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "X":
			break Loop
		default:
			convIn, err := strconv.Atoi(in)
			if err != nil {
				continue Loop
			} else {
				sli = append(sli, convIn)
				sort.Ints(sli)
				fmt.Println(sli)
			}
		}
	}
}
