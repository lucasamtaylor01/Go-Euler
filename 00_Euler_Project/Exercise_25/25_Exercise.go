/*
The Fibonacci sequence is defined by the recurrence relation:

F_n = F_{n - 1} + F_{n - 2}, where F_1 = 1 and F_2 = 1.

Hence the first 12 terms will be:

# EXEMPLE

The 12th term, F_{12}, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import (
	"fmt"
	"math/big"
)

// Function to count the digits of a bigInt
func DigitCounter(n *big.Int) int {
	str := n.String()
	return len(str)
}

func main() {
	//Inicializate the first terms of sequence
	i := big.NewInt(0)
	j := big.NewInt(1)

	//Define the counter and the limit
	var counter, limit int = 2, 1000

	for true {
		k := new(big.Int)
		k.Add(i, j)
		i = j
		j = k

		// Verify if j has 1000 digits
		if DigitCounter(j) == limit {

			// Show solution
			fmt.Println("Solution: ", counter)
			break
		}
		counter++
	}
}
