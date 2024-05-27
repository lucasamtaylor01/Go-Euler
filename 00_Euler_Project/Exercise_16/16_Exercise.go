// https://en.wikipedia.org/wiki/Digit_sum

package main

import (
	"fmt"
	"math"
)

func power2(n int) float64 {
	var base int = 1
	for i := 0; i < n; i++ {
		base *= 2
	}
	return float64(base)
}
func main() {
	var n float64 = power2(100)
	var k int = int(math.Log10(n))

	var sum int = 0
	for i := 0; i <= k; i++ {
		d_numerator := ((int(n) % int(math.Pow(10, float64(i+1)))) - (int(n) % int(math.Pow(10, float64(i)))))
		d_denominator := int(math.Pow(10, float64(i)))
		d := d_numerator / d_denominator
		sum += d
	}

	fmt.Println(sum)

}
