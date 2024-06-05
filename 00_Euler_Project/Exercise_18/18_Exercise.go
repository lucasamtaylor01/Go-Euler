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
}
