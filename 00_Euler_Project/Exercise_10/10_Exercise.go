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
	var sum int = 0
	for i := 2; i < 2000000; i++ {
		if IsPrime(i) {
			sum += i
		}
	}

	fmt.Println("Solution: ", sum)
}
