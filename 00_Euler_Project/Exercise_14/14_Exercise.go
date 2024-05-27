package main

import "fmt"

// inocent algorithm
func ChainLenght(n int) int {
	var chain []int

	for n != 1 {
		chain = append(chain, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3*n + 1)
		}
	}

	return len(chain) + 1
}

func main() {
	var biggest_chain, number_chain int = 0, 0
	for i := 2; i < 1000000; i++ {
		var chain_len int = ChainLenght(i)
		if biggest_chain < chain_len {
			biggest_chain = chain_len
			number_chain = i
		}
	}

	fmt.Println(number_chain)
}
