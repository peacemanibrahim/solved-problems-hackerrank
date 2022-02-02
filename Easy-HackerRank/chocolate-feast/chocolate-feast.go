// Link to problem statement:https://www.hackerrank.com/challenges/chocolate-feast/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                                   int32
	c                                   int32
	m                                   int32
	validity                            bool
	totalNumberOfChocolateBobbyHasEaten int32
)

func CheckDataValidity(n, c, m int32) bool {
	if n >= 2 && n <= 100000 && c >= 1 && c <= n && m >= 2 && m <= n {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func ChocolateFeast(n, c, m int32) int32 {
	validity = CheckDataValidity(n, c, m)
	if validity == true {
		noOfChocolateBobbyHasEatenInitially := n / c
		noOfWrappersLeft := noOfChocolateBobbyHasEatenInitially
		for noOfWrappersLeft >= m {
			noOfChocolateBobbyCanGetFromWrappers := noOfWrappersLeft / m
			wrappersLeftAfterTurningInWrappersForChocolate := noOfWrappersLeft % m
			if noOfChocolateBobbyCanGetFromWrappers >= 1 {
				noOfChocolateBobbyHasEatenInitially += noOfChocolateBobbyCanGetFromWrappers
				newlyEatenChocolateWrappers := noOfChocolateBobbyCanGetFromWrappers
				noOfWrappersLeft = newlyEatenChocolateWrappers + wrappersLeftAfterTurningInWrappersForChocolate
			}
		}
		totalNumberOfChocolateBobbyHasEaten = noOfChocolateBobbyHasEatenInitially
	}
	return totalNumberOfChocolateBobbyHasEaten
}

func main() {
	n = 6
	c = 2
	m = 2

	totalNumberOfChocolateBobbyHasEaten = ChocolateFeast(n, c, m)
	fmt.Println(totalNumberOfChocolateBobbyHasEaten)
}
