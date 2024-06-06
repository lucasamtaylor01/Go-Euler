package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("triangle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var m [][]int
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Fields(line)
		var row []int

		// Convert each string to int
		for _, str := range strNumbers {
			num, _ := strconv.Atoi(str)
			row = append(row, num)
		}
		m = append(m, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Try to create subslices to analyse the best path

	//var sum int = m[1][0] + m[2][0] + m[3][0]
	var i, j, k, l int = 0, 0, 0, 1
	// len(m)-2
	for l < 3 {
		fmt.Println(m[l][i], m[l+1][j], m[l+2][k])
		fmt.Println(l, ":", i, "/", l+1, ":", j, "/", l+2, ":", k)
		k++
		if k >= len(m[l]) {
			k = 0
			j++
			if j >= len(m[l]) {
				k = 0
				j = 0
				i++
				if i >= len(m[l]) {
					i = 0
					k = 0
					j = 0
					l++
				}
			}

		}

	}

}
