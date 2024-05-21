package main

import "fmt"

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
	fac_list, prime_list := FactorizationList(n)

	for i := 0; i < len(prime_list); i++ {
		var counter int = 1
		for j := 0; j < len(fac_list); j++ {
			if prime_list[i] == fac_list[j] {
				counter++
			}
		}
		divisor_list = append(divisor_list, counter)
	}

	var number_of_divisors int = 1
	for i := 0; i < len(divisor_list); i++ {
		number_of_divisors *= divisor_list[i]
	}

	return number_of_divisors
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

func TriangleNumber(n int) int {
	var triangle_number int = int(((n * n) + n) / 2)

	return triangle_number
}

func main() {
	var triangle_number, i int = 0, 1000
	for true {
		triangle_number = TriangleNumber(i)
		if DivisorCounter(TriangleNumber(i)) >= 500 {
			fmt.Println(i)
			break
		}
		i++
	}

	fmt.Println(triangle_number, DivisorCounter(triangle_number))
}
