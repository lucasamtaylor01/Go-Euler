/*A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 \times 99.
Find the largest palindrome made from the product of two 3-digit numbers.*/

package main

import "fmt"

func PalindromeVerify(n int) bool {
	var a1 int = n / 100000
	var a2 int = (n % 100000) / 10000
	var a3 int = ((n % 100000) % 10000) / 1000
	var a4 int = (((n % 100000) % 10000) % 1000) / 100
	var a5 int = ((((n % 100000) % 10000) % 1000) % 100) / 10
	var a6 int = (((((n % 100000) % 10000) % 1000) % 100) % 10)

	if a1 == a6 && a2 == a5 && a3 == a4 {
		return true
	} else {
		return false
	}
}

func main() {
	var a int = 999
	var candidate, solution int = 0, 0
	for i := a - 1; i >= 100; i-- {
		candidate = a * i
		if candidate/100000 != 0 {
			if PalindromeVerify(candidate) {
				if candidate > solution {
					solution = candidate
				}
			}
		} else {
			if i == 100 {
				a -= 1
				i = a - 1
			}
			continue
		}
	}
	fmt.Println("Solution: ", solution)
}
