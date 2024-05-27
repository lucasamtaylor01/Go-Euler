// https://en.wikipedia.org/wiki/Digit_sum

package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 = math.Pow(2, 1000)
	var k int = int(math.Log2(n))
	var sum float64 = 0
	for i := 0; i <= k; i++ {
		d_1 := int(n) % int(math.Pow(2, float64(i+1)))
		d_2 := int(n) % int(math.Pow(2, float64(i)))
		fmt.Println(d_1, d_2)
		num := d_1 - d_2
		den := int(math.Pow(2, float64(i)))
		sum += math.Abs(float64(num / den))
	}

	fmt.Println(sum)

}
