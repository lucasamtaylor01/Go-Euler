/*
Starting in the top left corner of a 2 x 2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.

	img

How many such routes are there through a 20 x 20 grid?
*/

package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	// Show combination
	fmt.Println(combin.Binomial(40, 20))
}
