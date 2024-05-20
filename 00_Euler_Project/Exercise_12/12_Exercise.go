package main

import (
	"fmt"
)

func NoContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}

func DivisorCounter(n int) int {
	divisor_list := make([]int, 0)
	prime_list, fac_list := FactorizationList(n)
	for i := 0; i < len(prime_list); i++ {
		var divisor_counter int = 1
		var j int = 0
		for j < len(fac_list) {
			if prime_list[i] == fac_list[j] {
				fmt.Println(prime_list[i], fac_list[j])
				divisor_list = append(divisor_list, divisor_counter)
				divisor_counter++
			} else {
				j++
			}
		}

	}

	fmt.Println(prime_list, fac_list, divisor_list)
	return 0
}
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
func main() {
	fmt.Println(DivisorCounter(48))
}

// 76576500
