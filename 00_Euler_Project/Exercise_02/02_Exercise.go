package main

import "fmt"

func main() {

	var i, j, k, sum int = 1, 2, 0, 2

	for k < 4000000 {
		k = i + j
		i = j
		j = k

		if k%2 == 0 {
			sum += k
		}

	}

	fmt.Println("Soma:", sum)
}
