/*The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143? */

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

	// Loop to search the largest prime factor of a
	for i := int(math.Sqrt(float64(600851475143))); i >= 1; i-- {
		if 600851475143%i == 0 {
			if IsPrime(i) {
				fmt.Println("Solution: ", i)
				break
			}
		}
	}

}
