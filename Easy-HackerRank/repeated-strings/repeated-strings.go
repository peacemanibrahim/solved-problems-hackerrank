// Link to problem statement: https://www.hackerrank.com/challenges/repeated-string/problem
package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	n                            int64
	validity                     bool
	numberOfAsInArrayOfS         int64
	additionalNumberOfAs         int64
	numberOfAsInInfiniteArrayOfS int64
)

func CheckDataValidity(s string, n int64) bool {
	arrayOfS := strings.Split(s, "")
	arrayLengthOfS := len(arrayOfS)
	basePower12 := int64(math.Pow(10, 12))
	if arrayLengthOfS >= 1 && arrayLengthOfS <= 100 && n >= 1 && n <= basePower12 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func RepeatedString(s string, n int64) int64 {
	validity = CheckDataValidity(s, n)
	if validity == true {
		arrayOfS := strings.Split(s, "")
		arrayLengthOfS := int64(len(arrayOfS))
		numberOfAsInArrayOfS = 0
		nDividedByArrayLengthOfS := n / arrayLengthOfS
		nDividedByArrayLengthOfSRemainder := n % arrayLengthOfS
		for i := range arrayOfS {
			if arrayOfS[i] == "a" {
				numberOfAsInArrayOfS++
			}
			if nDividedByArrayLengthOfSRemainder > 0 && i == int(nDividedByArrayLengthOfSRemainder)-1 {
				additionalNumberOfAs = numberOfAsInArrayOfS
			}
		}
		if nDividedByArrayLengthOfSRemainder == 0 {
			numberOfAsInInfiniteArrayOfS = nDividedByArrayLengthOfS * numberOfAsInArrayOfS
		} else {
			numberOfAsInInfiniteArrayOfS = nDividedByArrayLengthOfS * numberOfAsInArrayOfS
			numberOfAsInInfiniteArrayOfS += additionalNumberOfAs
		}
	}
	return numberOfAsInInfiniteArrayOfS
}

func main() {
	s := "aba"
	n = 10

	numberOfAsInInfiniteArrayOfS = RepeatedString(s, n)
	fmt.Println(numberOfAsInInfiniteArrayOfS)
}
