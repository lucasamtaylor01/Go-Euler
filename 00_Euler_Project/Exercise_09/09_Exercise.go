/* A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a^2 + b^2 = c^2.
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.Find the product abc.*/

package main

import "fmt"

func main() {
	// Indicator to stop the loop
	var ind int = 0

	// Loops to analyse candidates and define the solution
	for m := 21; ind == 0; m-- {
		for n := 0; n < m; n++ {
			if m*(m+n) == 500 {

				// Define the variables with Euclid's Formula
				var a, b, c int = (m * m) - (n * n), 2 * m * n, (m * m) + (n * n)

				// Show the solution
				fmt.Println("Solution: ", a*b*c)

				// End the Loops
				ind = 1
			}
		}
	}
}
