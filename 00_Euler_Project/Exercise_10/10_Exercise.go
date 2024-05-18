/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/

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
	var sum int = 0
	for i := 2; i < 2000000; i++ {
		if IsPrime(i) {
			sum += i
		}
	}

	fmt.Println("Solution: ", sum)
}
