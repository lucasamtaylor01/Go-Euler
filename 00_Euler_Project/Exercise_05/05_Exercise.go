/*2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible with no remainder by all of the numbers from 1 to 20?*/

package main

import "fmt"

// Verification if divisible 1-20
func DivibleVerify(n int) bool {
	for i := 1; i <= 20; i++ {
		if n%i == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

	// Inicializate de variables
	var candidate, solution int = 2520, 0

	// Loop of candidate verification
	for solution == 0 {
		if DivibleVerify(candidate) {
			solution = candidate
		} else {
			candidate += 1
		}
	}
	fmt.Println("Solution: ", solution)
}
