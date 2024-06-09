package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GreatPath(m [][]int, l int, last_index int) (int, int) {
	var sum, great, great_index int
	var floor_i, floor_j, floor_k int = last_index, 0, 0
	for i := floor_i - 1; i < floor_i+1; i++ {
		for j := floor_j; j < floor_j+2; j++ {
			for k := floor_k; k < floor_k+2; k++ {
				fmt.Println(m[l][i], m[l+1][j], m[l+2][k])
				sum += m[l][i] + m[l+1][j] + m[l+2][k]
			}
			if great < sum {
				great_index = i
				great = sum
			}
			sum = 0
			floor_k++
		}
		floor_j++
		floor_k--
		fmt.Println()
	}

	return m[l][great_index], great_index
}

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

	var great_path []int
	great_path = append(great_path, 75)
	great_path = append(great_path, 66)
	great_path = append(great_path, 98)
	var last_index, n int = 1, 0
	for i := 1; i < 14; i++ {
		n, last_index = GreatPath(m, i, last_index)
		if last_index == 0 {
			last_index++
		}
		great_path = append(great_path, n)
		fmt.Println(great_path, last_index)
	}

}
