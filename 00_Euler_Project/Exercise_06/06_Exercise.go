/*The sum of the squares of the first ten natural numbers is:
1^2 + 2^2 + ... + 10^2 = 385.

The square of the sum of the first ten natural numbers is:
(1 + 2 + ... + 10)^2 = 55^2 = 3025.

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.*/

package main

import "fmt"

func main() {
	// Initialize the variables
	var sum_of_squares, square_of_sum, solution int = 0, 0, 0

	// Add the acumulate values to variables
	for i := 1; i <= 100; i++ {
		sum_of_squares += (i * i)
		square_of_sum += i
	}

	// Define the difference
	solution = (square_of_sum * square_of_sum) - sum_of_squares

	// Show the solution
	fmt.Println("Solution: ", solution)
}
