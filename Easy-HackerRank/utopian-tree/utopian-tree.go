// Link to problem statement: https://www.hackerrank.com/challenges/utopian-tree/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                          int32
	validity                   bool
	defaultHeightOfUtopianTree int32
	finalHeightOfUtopianTree   int32
)

func CheckDataValidity(n int32) bool {
	if n >= 0 && n <= 60 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func CheckIfSpringOrSummerCycle(number int) string {
	var growthCycle string
	numberDividedBy2Remainder := number % 2
	if numberDividedBy2Remainder == 0 {
		growthCycle = "summer"
	} else {
		growthCycle = "spring"
	}
	return growthCycle
}

func UtopianTree(n int32) int32 {
	defaultHeightOfUtopianTree = 1
	if n == 0 {
		finalHeightOfUtopianTree = 1
		return finalHeightOfUtopianTree
	}
	validity = CheckDataValidity(n)
	if validity == true {
		for i := 1; i <= int(n); i++ {
			growthCycle := CheckIfSpringOrSummerCycle(i)
			if growthCycle == "spring" {
				defaultHeightOfUtopianTree *= 2
			} else if growthCycle == "summer" {
				defaultHeightOfUtopianTree += 1
			}
		}
		finalHeightOfUtopianTree = defaultHeightOfUtopianTree
	}
	return finalHeightOfUtopianTree
}

func main() {
	n = 5
	finalHeightOfUtopianTree = UtopianTree(n)
	fmt.Println(finalHeightOfUtopianTree)
}
