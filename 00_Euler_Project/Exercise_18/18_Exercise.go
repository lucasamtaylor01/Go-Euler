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

	var matrix [][]int
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
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var greatest_list []int
	greatest_list = append(greatest_list, 75)
	var pre_index = 0
	for i := 1; i < len(matrix); i++ {
		var great int = 0

		//Try to aply to next 3 steps
		if matrix[i][pre_index] >= matrix[i][pre_index+1] {
			great = matrix[i][pre_index]
		} else {
			great = matrix[i][pre_index+1]
			pre_index++
		}
		greatest_list = append(greatest_list, great)
	}

	fmt.Println(greatest_list)

	var sum int = 0
	for i := 0; i < len(greatest_list); i++ {
		sum += greatest_list[i]
	}
	fmt.Println(sum)
}
