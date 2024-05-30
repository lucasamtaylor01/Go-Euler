/*
n! means n x (n - 1) x ... x 3 x 2 x 1.
For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Source: https://stackoverflow.com/qustions/11270547/go-big-int-factorial-with-recursion
func factorial(n *big.Int) (result *big.Int) {
	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
}

func main() {
	// Define 100!
	a := factorial(big.NewInt(100))

	// Convert the big Int to string
	str := a.String()

	// Define the sum
	var sum int = 0

	// Loop to sum the digits of 100!
	for i := 0; i < len(str); i++ {
		n, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Print(err)
		} else {
			sum += n
		}
	}

	// Show solution
	fmt.Println(sum)
}
