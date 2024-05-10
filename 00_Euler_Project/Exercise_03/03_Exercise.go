/*The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143? */

package main

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var a int = 600851475143
	for i := int(math.Sqrt(float64(a))); i >= 1; i-- {
		if a%i == 0 {
			if IsPrime(i) {
				fmt.Println("Solution: ", i)
				break
			}
		}
	}

}
