/*
The following iterative sequence is defined for the set of positive integers:

n -> n/2 (n is even)
n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1.

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/
package main

import "fmt"

// Verify the chain lenght
func ChainLenght(n int) int {
	var chain_counter int = 1
	for n != 1 {
		chain_counter++

		// Apply collatz number logic
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3*n + 1)
		}
	}

	return chain_counter
}

func main() {
	// Create variables to analyse the candidates
	var biggest_chain, number_chain int = 0, 0

	// Loop to analyse candidates under 10^6
	for i := 2; i < 1000000; i++ {
		var chain_len int = ChainLenght(i)
		if biggest_chain < chain_len {
			biggest_chain = chain_len
			number_chain = i
		}
	}

	// Show solution
	fmt.Println(number_chain)
}

//I know that is not efficient. In future, rewrite the code can be a good ideia

