# What I Learned from Solving the 18th Problem

# Matrix Idea
My first ideia was create a matrix to organize the numbers. I tried search online, but I didn't find a good text explaining the concept. So, I used CHATGPT. Let's analyse the code and learn how it works.

The full code generated is:
```go
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

```

## Questions

1. What's `defer`?

    Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. In this case, `defer` is used to close the file .txt as soon function main ends.

    *Source:* https://go.dev/doc/effective_go#defer

2. What's `bufio`?  
    
    `bufio` is a Go package. Package `bufio`  implements buffered I/O. 

    *Source:* https://pkg.go.dev/bufio

3. And why we use it?

    This question is related with buffered. Buffered is a optimum way to read a file. I like the way that CHATGPT explain that to me:

    Imagine that you are on a library and take a book. 

    - **No buffer**: You read one word at a time and put the book back on the shelf after each word. This would be very inefficient because you spend a lot of time picking up and putting away the book.

    - **With buffer**: You take the book from the shelf, read a whole chapter in one go and then put the book away. This is much more efficient because you minimise the time spent picking up and putting away the book and can concentrate more on reading the content.

4. About ` for scanner.Scan()`:

    - To understand `scanner`, we need to analyse this line *again*:
        ```go
        scanner := bufio.NewScanner(file)
        ```

        `scanner` is a estructure `struct` of type `bufio.Scanner` that alow read file's data using buffer explained on question 3.
    
        *Source:* https://pkg.go.dev/bufio#Scanner
    
    - `Scan()` isn't a function, but is a **method**. A method is a function with a special receiver argument.

        *Source:* https://go.dev/tour/methods/1

    - How `Scan()`works?
        > Scan advances the Scanner to the next token, which will then be available through the `Scanner.Bytes` or `Scanner.Text` method. It returns `false` when there are no more tokens (...)

        *Source:* https://pkg.go.dev/bufio#Scanner.Scan
    
5. How `strings.Fields(line)` works? 
    
    > Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space. 

    *Source:* https://pkg.go.dev/strings#Fields

    In other words:

    - First, the function receive a string
    - After, the function divides the strings in substrings according to blank spaces
    - The function returns a slice with substrings

    *Exemple:* https://go.dev/play/p/i_3x-wRvAxx

6. What's `_` on `for _, str := range strNumbers`?

    *Source:* https://go.dev/doc/effective_go#blank

    
