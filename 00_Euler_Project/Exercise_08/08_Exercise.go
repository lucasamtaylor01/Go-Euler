package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("number.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	var number_list []int
	for i := 0; i < len(b); i++ {
		n, err := strconv.Atoi(string(b[i]))
		if err != nil {
			fmt.Print("ué")
		} else {
			number_list = append(number_list, n)
		}

	}

	for i := 0; i < len(number_list); i++ {
		fmt.Print(number_list[i])
	}
}
