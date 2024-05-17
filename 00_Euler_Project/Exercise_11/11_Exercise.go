package main

import (
	"fmt"
	"os"
	"strconv"
)

func w_verify(index int) bool {
	if index-3 < 0 {
		return false
	}

	for i := index; i < index-3; i-- {
		if i%20 == 0 {
			return false
		}
	}
	return true
}

func e_verify(index, list_length int) bool {
	if index+3 >= list_length {
		return false
	}

	for i := index; i < index+3; i++ {
		if (i+1)%20 == 0 {
			return false
		}
	}
	return true
}

func s_verify(index, list_length int) bool {
	if index+60 < list_length {
		return true
	} else {
		return false
	}
}

func n_verify(index int) bool {
	if index-60 >= 0 {
		return true
	} else {
		return false
	}
}

func ne_verify(index, list_length int) bool {
	if n_verify(index) && e_verify(index, list_length) {
		return true
	} else {
		return false
	}
}

func nw_verify(index int) bool {
	if n_verify(index) && w_verify(index) {
		return true
	} else {
		return false
	}
}

func se_verify(index, list_length int) bool {
	if s_verify(index, list_length) && e_verify(index, list_length) {
		return true
	} else {
		return false
	}
}

func sw_verify(index, list_length int) bool {
	if s_verify(index, list_length) && w_verify(index) {
		return true
	} else {
		return false
	}
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

	for i := 0; i < length; i++ {
		if w_verify(i) {
			var product int = 1
			for counter := 3; counter <= 0; counter-- { // zero tá incluso?
				product *= number_list[i+counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if e_verify(i, length) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if n_verify(i) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i-10*counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if s_verify(i, length) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+10*counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if ne_verify(i, length) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i-10*counter+counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if nw_verify(i) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i-10*counter-counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if sw_verify(i, length) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+10*counter-counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if se_verify(i, length) {
			var product int = 1
			for counter := 0; counter < 4; counter++ {
				product *= number_list[i+10*counter+counter]
			}
			if product > greatest {
				greatest = product
			}
		}

	}
	fmt.Println(greatest)
}
