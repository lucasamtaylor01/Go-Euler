package main

import "fmt"

// Verify if the number search isn't on the list
// Source: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func NoContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}

// Function to count how many divisors a number has
func DivisorCounter(n int) int {

	// Create a variable to store the divisors
	divisor_list := make([]int, 0)

	// Variables that store the factorization of the number and the prime number that composes the number
	fac_list, prime_list := FactorizationList(n)

	// Loop to count de divisors and store them on a slice
	for i := 0; i < len(prime_list); i++ {
		var counter int = 1
		for j := 0; j < len(fac_list); j++ {
			if prime_list[i] == fac_list[j] {
				counter++
			}
		}
		divisor_list = append(divisor_list, counter)
	}

	// Define the number of divisors
	// Source: https://mathschallenge.net/library/number/number_of_divisors
	var number_of_divisors int = 1
	for i := 0; i < len(divisor_list); i++ {
		number_of_divisors *= divisor_list[i]
	}

	return number_of_divisors
}

// Function that return a slice with factorization and prime that composes the number
func FactorizationList(n int) ([]int, []int) {
	fac_list := make([]int, 0)
	prime_list := make([]int, 0)
	var i int = 2
	for n != 1 {
		if n%i == 0 {
			fac_list = append(fac_list, i)
			if NoContains(prime_list, i) {
				prime_list = append(prime_list, i)
			}
			n = n / i
		} else {
			i++
		}
	}

	return fac_list, prime_list
}

// Function to generate a triangle number
// Source: https://en.wikipedia.org/wiki/Triangular_number
func TriangleNumber(n int) int {
	var triangle_number int = int(((n * n) + n) / 2)

	return triangle_number
}

func main() {
	var triangle_number, i int = 0, 1
	// Loop to search the first triangle number to have over five hundred divisors
	for true {
		triangle_number = TriangleNumber(i)
		if DivisorCounter(TriangleNumber(i)) >= 500 {
			break
		}
		i++
	}

	fmt.Println(triangle_number, DivisorCounter(triangle_number))
}
