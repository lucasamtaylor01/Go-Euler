/*
The series, 1^1 + 2^2 + 3^3 + ... + 10^{10} = 10405071317.
Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^{1000}.
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	// Inicializate the variables
	j := big.NewInt(1)
	sum := big.NewInt(0)

	// Loop to sum self power
	for counter := 1; counter < 1000; counter++ {

		// Define self power
		k := new(big.Int)
		k = k.Exp(j, j, nil)

		sum = sum.Add(sum, k)
		j = j.Add(j, big.NewInt(1))
	}

	// Convert to string to deal with digits
	str := sum.String()

	// Loop to show only the ten last digits
	for i := len(str) - 10; i < len(str); i++ {
		n, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Printf("Erro")
		}

		print(n)
	}
	fmt.Println()
}
