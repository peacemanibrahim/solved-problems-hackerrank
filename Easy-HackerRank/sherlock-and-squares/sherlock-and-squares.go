// Link to problem statement: https://www.hackerrank.com/challenges/sherlock-and-squares/problem
package main

import (
	"fmt"
	"math"
)

var (
	a                   int32
	b                   int32
	validity            bool
	squareIntegersCount int32
)

func CheckDataValidity(a, b int32) bool {
	basePower9 := int32(math.Pow(10, 9))
	if a >= 1 && b >= a && b <= basePower9 {
		validity = true
	} else {
		fmt.Println("Invalid data for a or b. Check the data.")
	}
	return validity
}

func Squares(a, b int32) int32 {
	squareIntegersCount = 0
	validity = CheckDataValidity(a, b)
	if validity == true {
		for i := 1; i <= 31622; i++ {
			squareOfI := int32(math.Pow(float64(i), 2))
			if squareOfI >= a && squareOfI <= b {
				squareIntegersCount++
			}
			if squareOfI >= b {
				break
			}
		}
	}
	return squareIntegersCount
}

func main() {
	a = 17
	b = 24

	squareIntegersCount = Squares(a, b)
	fmt.Println(squareIntegersCount)
}

// Note: 31622 is the highest square "root" in that is lower than 1 billion (1000000000).
