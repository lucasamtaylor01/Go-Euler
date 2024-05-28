package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.Exp(big.NewInt(2), big.NewInt(1000), nil)

	fmt.Println(a)
}

// https://pkg.go.dev/math/big
