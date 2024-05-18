package main

import (
	"fmt"
)

func DivisorCounter(n int) int {
	var counter int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			counter++
		}
	}

	return counter
}

func main() {

	fmt.Println(DivisorCounter(28))
}
