package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	a := new(big.Int)
	a.Exp(big.NewInt(2), big.NewInt(1000), nil)
	str := a.String()

	var sum int = 0
	for i := 0; i < len(str); i++ {
		n, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Print(err)
		} else {
			sum += n
		}
	}

	fmt.Println(sum)
}
