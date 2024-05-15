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
	if index+30 < list_length {
		return true
	} else {
		return false
	}
}

func n_verify(index int) bool {
	if index-30 >= 0 {
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
	var greatest, lenght int = 0, 5

	for i := 0; i < lenght; i++ {
		if w_verify(i) {
			var product int = 1
			for j := i - 3; j <= i; j++ {
				product *= number_list[j]
			}
			if product > greatest {
				greatest = product
			}
		}

		if e_verify(i, lenght) {
			var product int = 1
			for j := i; j <= i+3; j++ {
				product *= number_list[j]
			}
			if product > greatest {
				greatest = product
			}
		}

		if n_verify(i) {
			var product, counter int = 1, 1
			for j := i; counter < 4; j++ {
				product *= number_list[j-10*counter]
			}
			if product > greatest {
				greatest = product
			}
		}

		if s_verify(i, lenght) {
			var product, counter int = 1, 0
			for j := i; counter < 4; j++ {
				product *= number_list[j+10*counter]
			}
			if product > greatest {
				greatest = product
			}
		}
	}
	fmt.Println(greatest)
}


/* Lembrar que é uma matriz 20x20 e para verificar que existe uma diagonal é necessário as condições de norte e sul, junto com a verificação de que é menor que 20 ou não 
logaicamente, dependendo da casa atual*/
