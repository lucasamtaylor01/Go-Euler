package main

import (
	"fmt"
	"os"
	"strconv"
)

func horizontal_verify(index int, length int) bool {
	if index+3 > length-1 {
		return false
	}

	if (index+4)%20 == 0 {
		return false
	}

	return true
}

func vertical_verify(index int, length int) bool {
	if index+60 > length-1 {
		return false
	}
	return true
}

func down_diagonal_verify(index int, length int) bool {
	if index+60+3 > length-1 {
		return false
	}
	return true
}

func up_diagonal_verify(index int) bool {
	if index-60-3 <= 0 {
		return false
	}
	return true
}

func main() {

	//Read file
	b, err := os.ReadFile("20grid.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Create a var slice
	var number_list []int

	// Convertion: bytes -> str -> int
	for i := 0; i < len(b)-1; i++ { // len b?
		if b[i] >= 48 && b[i] <= 57 && b[i+1] >= 48 && b[i+1] <= 57 {
			c := string(b[i]) + string(b[i+1])
			n, err := strconv.Atoi(c)
			if err != nil {
				fmt.Print(err)
			} else {
				number_list = append(number_list, n)
			}
		}
	}
	var greatest, length int = 0, len(number_list)
	var great_list [4]int

	for i := 0; i < length; i++ {
		if horizontal_verify(i, length) {
			var product int = 1
			var candidate [4]int
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+counter]
				candidate[counter] = number_list[i+counter]
			}

			if product > greatest {
				greatest = product
				for i := 0; i < 4; i++ {
					great_list[i] = candidate[i]
				}
			}
		}

		if vertical_verify(i, length) {
			var product int = 1
			var candidate [4]int
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+20*counter]
				candidate[counter] = number_list[i+20*counter]
			}

			if product > greatest {
				greatest = product
				for i := 0; i < 4; i++ {
					great_list[i] = candidate[i]
				}
			}
		}

		if down_diagonal_verify(i, length) {
			var product int = 1
			var candidate [4]int
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+20*counter-counter]
				candidate[counter] = number_list[i+20*counter+counter]
			}
			fmt.Println(candidate)
			if product > greatest {
				greatest = product
				for i := 0; i < 4; i++ {
					great_list[i] = candidate[i]
				}
			}
		}

		if up_diagonal_verify(i) {
			var product int = 1
			var candidate [4]int
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i-20*counter-counter]
				candidate[counter] = number_list[i-20*counter-counter]
			}
			 // fmt.Println(candidate)
			if product > greatest {
				greatest = product
				for i := 0; i < 4; i++ {
					great_list[i] = candidate[i]
				}
			}
		}
	}
	fmt.Println(greatest)
	fmt.Println(great_list)

}
/* 70600674
[89 69 7 46]
/*
