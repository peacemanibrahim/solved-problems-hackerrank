package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := "799898013447209990576845913871859165418584121624031762316631"
	y := "6046724675206781253805417697364063874537349942332939012160502967377173420319837265456476"
	// Line 12 and 13 below is on converting strings to big integers
	bigIntRepOfNumberStringX, _ := new(big.Int).SetString(x, 10)
	bigIntRepOfNumberStringY, _ := new(big.Int).SetString(y, 10)
	// Line 16 below is on checking which big integer is greater in x and y. If answer is 1, then x is greater.
	// If answer is -1, then y is greater. If answer is 0 then they are the same.
	answer := bigIntRepOfNumberStringX.Cmp(bigIntRepOfNumberStringY)
	fmt.Println(answer)
}
