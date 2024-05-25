package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fill the slice with zeros
func FillSlice(a []int) []int {
	for i := 0; i < 52; i++ {
		a = append(a, 0)
	}

	return a
}

// Defining hundreds, tens and units of a number
func NumberDistribution(n int) (int, int, int) {
	var a int = n / 100
	var b int = (n % 100) / 10
	var c int = ((n % 100) % 10)

	return a, b, c
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

	// Variable that stores the sum of the column
	var partial_sum int = 0

	// Creating the variable that stores the 52-digit numbers
	var big_number []int
	big_number = FillSlice([]int(big_number))

	// Loop control variables
	var a, c int = len(number_list) - 1, 49
	var index int = len(big_number) - 1
	var i int = c

	for true {

		// Sum of columns
		partial_sum += number_list[i]
		if i == a {
			var x, y, z int = NumberDistribution(partial_sum)

			// Distribution of hundreds, tens and units of a number on slice
			big_number[index] += z
			if big_number[index] >= 10 {
				big_number[index] -= 10
				big_number[index-1]++
			}
			big_number[index-1] += y
			if big_number[index-1] >= 10 {
				big_number[index-1] -= 10
				big_number[index-2]++
			}
			big_number[index-2] += x
			if big_number[index-2] >= 10 {
				big_number[index-2] -= 10
				big_number[index-3]++
			}

			// Loop control
			partial_sum = 0
			index--
			a--
			c--
			i = c

			if i < 0 {
				break
			}

		} else {
			i += 50
		}

	}

	//Show answer
	for i := 0; i < 10; i++ {
		print(big_number[i])
	}

}
