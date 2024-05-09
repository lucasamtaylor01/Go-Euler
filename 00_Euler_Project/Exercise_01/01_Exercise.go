/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get and 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
    // Initialize sum to accumulate multiples of 3 or 5.
    sum := 0
    
    // Loop through numbers from 0 to 999.
    for i := 0; i < 1000; i++ {
        // Add to sum if i is a multiple of 3 or 5.
        if i%5 == 0 || i%3 == 0 {
            sum += i
        }
    }
    
    // Print the computed sum.
    fmt.Println(sum)
}
