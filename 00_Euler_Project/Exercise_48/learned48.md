# What I Learned from Solving the 48th Problem

I think the hardest part was converting to a string and extracting only the last ten numbers of the sum. It is not the first time I have dealt with this, but it is still difficult for me.


1. **Converting a BigInt to a String or Integer in Go**
    
    It's easy and not my first time doing it. Here is the code that helped me:
    ```go
    bigint := big.NewInt(123)
    bigstr := bigint.String()
    ```
    Source: https://stackoverflow.com/questions/11810948/convert-a-bigint-to-a-string-or-integer-in-go

2. My first attempt had a mistake:
    ```go
    n, err := strconv.Atoi(str[i])
    ```
    The correct way to do this is:
    ```go
    n, err := strconv.Atoi(string(str[i]))
    ```

    - `str[i]` returns a byte
    - `string(str[i])` converts this byte into a string
    - `strconv.Atoi` expects a string as input, not a byte.
