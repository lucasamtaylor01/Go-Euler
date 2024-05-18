/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get and 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
	// Initialize sum to accumulate the multiples of 3 and 5
	sum := 0

	// Loop to search and add the multiples of 3 and 5
	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}

	// Show solution
	fmt.Print(sum)
}
