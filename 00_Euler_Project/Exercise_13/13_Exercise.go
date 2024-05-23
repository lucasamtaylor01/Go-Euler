package main

import (
	"fmt"
	"os"
	"strconv"
)

func FillSlice(a []int) []int {
	for i := 0; i < 50; i++ {
		a = append(a, 0)
	}

	return a
}

func main() {

	// Read file
	b, err := os.ReadFile("50digits.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Create a var slice
	var number_list []int

	// Convertion: bytes -> str -> int
	for i := 0; i < len(b); i++ {
		if string(b[i]) != "\n" {
			n, err := strconv.Atoi(string(b[i]))
			if err != nil {
				fmt.Print(err)
			} else {
				number_list = append(number_list, n)
			}
		}
	}

	var partial_sum, accumulated int = 0, 0
	var a, c int = len(number_list) - 1, 49
	var big_number []int

	big_number = FillSlice([]int(big_number))

	for i := c; i <= a; i += 50 {
		partial_sum += number_list[i]
		if partial_sum > 10 {
			accumulated++
			partial_sum -= 10
		}
		if i == a {
			fmt.Println(partial_sum, accumulated)
			fmt.Println(accumulated / 10)
			break
		}

	}
}
