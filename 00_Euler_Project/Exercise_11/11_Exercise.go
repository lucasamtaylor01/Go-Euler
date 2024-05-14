package main

import (
	"fmt"
	"os"
	"strconv"
)

func w_verify(index int) bool {
	if index-3 >= 0 {
		return true
	} else {
		return false
	}
}

func e_verify(index, list_length int) bool {
	if index+3 < list_length {
		return true
	} else {
		return false
	}
}

func s_verify(index, list_length int) bool {
	if index+40 < list_length {
		return true
	} else {
		return false
	}
}

func n_verify(index int) bool {
	if index-40 >= 0 {
		return true
	} else {
		return false
	}
}

func ne_verify(index, list_lenght int) bool {
	if n_verify(index) && e_verify(index, list_lenght) {
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

	fmt.Println(len(number_list))

}
