package main

import (
	"fmt"
	"os"
)

func main() {
	// Read file
	b, err := os.ReadFile("50digits.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Convertion: bytes -> str
	var j int = 0
	c := make([]string, 50)
	for i := 0; i < len(b)-1; i++ {
		if b[i] == 10 {
			j = 0
			fmt.Println(c)
		} else {
			c[j] = string(b[i])
			j++
		}
	}

}
