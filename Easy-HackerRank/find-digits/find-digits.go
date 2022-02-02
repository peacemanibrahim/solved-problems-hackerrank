// Link to problem statement: https://www.hackerrank.com/challenges/find-digits/problem
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	n            int32
	validity     bool
	divisorCount int32
)

func CheckDataValidity(n int32) bool {
	basePower9 := int32(math.Pow(10, 9))
	if n > 0 && n < basePower9 {
		validity = true
	} else {
		log.Fatal("Invalid data for n. Check the data.")
	}
	return validity
}

func FindDigits(n int32) int32 {
	validity = CheckDataValidity(n)
	divisorCount = 0
	if validity == true {
		stringRepresentationOfNumber := strconv.Itoa(int(n))
		arrayRepresentationOfNumber := strings.Split(stringRepresentationOfNumber, "")
		for i := range arrayRepresentationOfNumber {
			eachDigitInArray := arrayRepresentationOfNumber[i]
			integerReprOfEachDigitInArray, _ := strconv.Atoi(eachDigitInArray)
			if integerReprOfEachDigitInArray != 0 {
				nDividedByDigitRemainder := n % int32(integerReprOfEachDigitInArray)
				if nDividedByDigitRemainder == 0 {
					divisorCount++
				}
			}
		}
	}
	return divisorCount
}

func main() {
	n = 1012

	divisorCount = FindDigits(n)
	fmt.Println(divisorCount)
}
