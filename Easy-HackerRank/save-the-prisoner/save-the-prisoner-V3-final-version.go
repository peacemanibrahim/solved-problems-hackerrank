// Link to problem statement: https://www.hackerrank.com/challenges/save-the-prisoner/problem
// This is the solution to the problem using dynamic programming. It takes lesser time to run.
// All test cases passed on hackerrank
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	n                              int32
	m                              int32
	s                              int32
	validity                       bool
	theChairNoOfPrisonerToBeWarned int32
)

func CheckDataValidity(n, m, s int32) bool {
	basePower9 := int32(math.Pow(10, 9))
	if n >= 1 && n <= basePower9 && m >= 1 && m <= basePower9 && s >= 1 && s <= n {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func SaveThePrisoner(n, m, s int32) int32 {
	validity = CheckDataValidity(n, m, s)
	if validity == true {
		mPlusSResultMinus1 := (m + s) - 1
		if mPlusSResultMinus1 <= n {
			theChairNoOfPrisonerToBeWarned = mPlusSResultMinus1
		} else if mPlusSResultMinus1 > n {
			divisionResult := (mPlusSResultMinus1 / n)
			valueToBeSubtracted := (divisionResult * n)
			theChairNoOfPrisonerToBeWarned = (mPlusSResultMinus1 - valueToBeSubtracted)
			if theChairNoOfPrisonerToBeWarned == 0 {
				theChairNoOfPrisonerToBeWarned += n
			}
		}
	}
	return theChairNoOfPrisonerToBeWarned
}

func main() {
	n = 7
	m = 19
	s = 2
	theChairNoOfPrisonerToBeWarned = SaveThePrisoner(n, m, s)
	fmt.Println(theChairNoOfPrisonerToBeWarned)
}
