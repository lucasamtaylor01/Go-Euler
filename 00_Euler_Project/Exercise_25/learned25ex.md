# WHAT I LEARNED

## CONVERTION BIGINT TO STRING
- https://stackoverflow.com/questions/11810948/convert-a-bigint-to-a-string-or-integer-in-go
- We can use: func (*Int) String 
- Source: golang.org/pkg/math/big/#Int.String
- How I used: `str := n.String()`


## FUNCTION WITH BIG INT AS A PARAMETER
- Using a pointer (*big.Int) in the DigitCounter function avoids large copies, making the code more efficient;
- Always use pointers to bigint, for example:
```
func AddBigInts(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}
```
