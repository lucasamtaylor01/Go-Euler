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

	var l int = 1
	var ceil_i, ceil_j, ceil_k int = len(m[l]), len(m[l]), len(m[l+1])
	var floor_i, floor_j, floor_k int = 0, 0, 0
	for i := floor_i; i < ceil_i; i++ {
		for j := floor_j; j < ceil_j; j++ {
			for k := floor_k; k < floor_k+2; k++ {
				fmt.Println(m[l][i], m[l+1][j], m[l+2][k])
				fmt.Println(l, ":", i, "/", l+1, ":", j, "/", l+2, ":", k)
				fmt.Println(floor_k, ":", ceil_k)
			}
			floor_k++
		}
		floor_j++
		ceil_j++
	}
	floor_i++
	ceil_i++
	l++

}

/*
	95-17-18
	95-17-35
	95-47-18
	95-47-35
	**64-17-18
	64-17-35
	64-47-
*/
