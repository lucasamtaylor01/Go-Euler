package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Abre o arquivo
	file, err := os.Open("triangle.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)

	// Lê o arquivo linha por linha
	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Fields(line)
		var row []int

		// Converte cada string para int
		for _, str := range strNumbers {
			num, _ := strconv.Atoi(str)
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	fmt.Println(matrix[1][0])

}
