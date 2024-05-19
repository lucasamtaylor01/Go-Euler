package main

import "fmt"

func DivisorCounter(n int) int {
	var counter int = 1
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			counter++
		}
	}

	return counter
}

func main() {
	var triangle_number, great, i int = 0, 0, 1
	for true {
		triangle_number += i
		if DivisorCounter(triangle_number) > great {
			fmt.Println(triangle_number, DivisorCounter(triangle_number), i, great)
			great = DivisorCounter(triangle_number)
		}

		if DivisorCounter(triangle_number) >= 500 {
			break
		}

		i++
	}
}
