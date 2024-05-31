# WHAT I LEARNED

# CONVERTION BIGINT TO STRING
https://stackoverflow.com/questions/11810948/convert-a-bigint-to-a-string-or-integer-in-go

- We can use: func (*Int) String 
- Source: golang.org/pkg/math/big/#Int.String
- How I used: `str := n.String()`

# FUNCTION WITH BIG INT AS A PARAMETER
- Usar um ponteiro (*big.Int) na função DigitCounter evita cópias grandes, tornando o código mais eficiente;
- Sempre usar ponteiros para bigint, exemplo:
```
func AddBigInts(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}
```
