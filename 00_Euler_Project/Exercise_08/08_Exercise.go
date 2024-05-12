package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//Read file
	b, err := os.ReadFile("number.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	// Create a var slice
	var number_list []int

	// Convertion: bytes -> str -> int
	for i := 0; i < len(b); i++ {
		n, err := strconv.Atoi(string(b[i]))
		if err != nil {
			fmt.Print(err)
		} else {
			number_list = append(number_list, n)
		}

	}

	// Show the int values
	for i := 0; i < len(number_list); i++ {
		fmt.Print(number_list[i])
	}
}
