package main

import (
	"fmt"
	"os"
	"strconv"
)

func FillSlice(a []int) []int {
	for i := 0; i < 52; i++ {
		a = append(a, 0)
	}

	return a
}

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

	var partial_sum int = 0
	var a, c int = len(number_list) - 1, 49
	var big_number []int

	big_number = FillSlice([]int(big_number))
	var index int = len(big_number) - 1
	var i int = c

	for true {
		partial_sum += number_list[i]
		if i == a {
			var x, y, z int = NumberDistribution(partial_sum)
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

	for i := 0; i < 10; i++ {
		print(big_number[i])
	}

}
