// This is the solution to the problem using for loop. It takes more time to run with large input values
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
	if n == 1 {
		return 1
	}
	if m == 1 {
		return s
	}
	validity = CheckDataValidity(n, m, s)
	if validity == true {
		for i := 2; i <= int(m); i++ {
			if s == n {
				s = 0
			}
			s++
		}
		theChairNoOfPrisonerToBeWarned = s
	}
	return theChairNoOfPrisonerToBeWarned
}

func main() {
	n = 4
	m = 6
	s = 2
	theChairNoOfPrisonerToBeWarned = SaveThePrisoner(n, m, s)
	fmt.Println(theChairNoOfPrisonerToBeWarned)
}

// Note: n is the number of prisoners
// Note: m is the number of candies
// Note: s is the chair number to begin passing out candies from
