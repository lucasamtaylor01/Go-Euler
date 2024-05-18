package main

import (
	"fmt"
	"os"
	"strconv"
)

// Checking if the current element has three horizontally adjacent elements
func HorizontalVerify(index int, length int) bool {
	if index+3 > length-1 {
		return false
	}

	if (index+4)%20 == 0 {
		return false
	}

	return true
}

// Checking if the current element has three vertically adjacent elements
func VerticalVerify(index int, length int) bool {
	if index+61 > length {
		return false
	}
	return true
}

// Checking if the current element has three diagonal (NW-SE) adjacent elements
func DiagonalTypeOneVerify(index int, length int) bool {
	if index+64 > length {
		return false
	}
	return true
}

// Checking if the current element has three diagonal (SW-NE) adjacent elements
func DiagonalTypeTwoVerify(index int) bool {
	if index-63 <= 0 {
		return false
	}
	return true
}

// Update the great_list if there's a bigger product
func UpdateGreatestList(great_list [4]int, candidate [4]int) [4]int {
	for i := 0; i < 4; i++ {
		great_list[i] = candidate[i]
	}

	return great_list
}

func main() {

	// Read file
	b, err := os.ReadFile("20grid.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Create a var slice
	var number_list []int

	// Convertion: bytes -> str -> int
	for i := 0; i < len(b)-1; i++ {
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

	// Creation of variables to store the greatest product and the four elements that originate it
	var greatest int = 0
	var great_list [4]int

	// Loop to search on the grid the greatest product
	for i := 0; i < len(number_list); i++ {
		if HorizontalVerify(i, len(number_list)) {

			// Creation of variables to analyse the candidates to greatest product
			var product int = 1
			var candidate [4]int

			// Loop to search for candidates in the grid horizontals
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+counter]
				candidate[counter] = number_list[i+counter]
			}

			// Update the greatest and great_list, if it's necessary
			if greatest < product {
				greatest = product
				great_list = UpdateGreatestList(great_list, candidate)
			}
		}

		if VerticalVerify(i, len(number_list)) {

			// Creation of variables to analyse the candidates to greatest product
			var product int = 1
			var candidate [4]int

			// Loop to search for candidates in the grid verticals
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+20*counter]
				candidate[counter] = number_list[i+20*counter]
			}

			// Update the greatest and great_list, if it's necessary
			if greatest < product {
				greatest = product
				great_list = UpdateGreatestList(great_list, candidate)
			}
		}

		if DiagonalTypeOneVerify(i, len(number_list)) {

			// Creation of variables to analyse the candidates to greatest product
			var product int = 1
			var candidate [4]int

			// Loop to search for candidates on the diagonals (NW-SE) of the grid
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+20*counter-counter]
				candidate[counter] = number_list[i+20*counter+counter]
			}

			// Update the greatest and great_list, if it's necessary
			if greatest < product {
				greatest = product
				great_list = UpdateGreatestList(great_list, candidate)
			}
		}

		if DiagonalTypeTwoVerify(i) {

			// Creation of variables to analyse the candidates to greatest product
			var product int = 1
			var candidate [4]int

			// Loop to search for candidates on the diagonals (SW-NE) of the grid
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i-20*counter-counter]
				candidate[counter] = number_list[i-20*counter-counter]
			}

			// Update the greatest and great_list, if it's necessary
			if greatest < product {
				greatest = product
				great_list = UpdateGreatestList(great_list, candidate)
			}
		}
	}

	// Show results
	fmt.Println(greatest)
	fmt.Println(great_list)

}
