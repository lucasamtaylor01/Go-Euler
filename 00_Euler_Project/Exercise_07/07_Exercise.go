/* By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number? */

package main

import (
	"fmt"
	"math"
)

// Verify if the number is a prime
func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Initialize the variables
	var candidate, counter int = 2, 0

	//Loop to search the 10001st prime
	for counter < 10001 {
		if IsPrime(candidate) {
			counter++
		}
		candidate++
	}

	// Show the solution and correct the answer
	fmt.Println("Solution: ", candidate-1)

}
