// Link to problem statement: https://www.hackerrank.com/challenges/drawing-book/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	n                      int32
	p                      int32
	validity               bool
	pages                  int32
	defaultPage            int32
	turnedPagesFromFront   int32
	turnedPagesFromBack    int32
	minimumNoOfPagesTurned int32
)

func CheckDataValidity(n, p int32) bool {
	basePower5 := math.Pow(10, 5)
	if n >= 1 && n <= int32(basePower5) && p >= 1 && p <= n {
		validity = true
	} else {
		log.Fatal("Invalid data for n and/or p. Check the data")
	}
	return validity
}

func CheckIfNumberIsEvenOrOdd(number int32) string {
	var answer string
	numberDividedBy2Remainder := number % 2
	if numberDividedBy2Remainder == 0 {
		answer = "even number"
	} else {
		answer = "odd number"
	}
	return answer
}

func FindNoOfTurnedPagesFromTheFront(n, p int32) int32 {
	defaultPage = 1
	pages = defaultPage
	turnedPagesFromFront = 0
	shouldIBreakTheLoop := ""
	for pages <= n+1 {
		if p == pages || p == pages-1 {
			shouldIBreakTheLoop = "yes"
		}
		if shouldIBreakTheLoop == "yes" {
			break
		}
		turnedPagesFromFront += 1
		pages += 2
	}
	return turnedPagesFromFront
}

func FindNoOfTurnedPagesFromTheBack(n, p int32) int32 {
	number := CheckIfNumberIsEvenOrOdd(n)
	if number == "even number" {
		defaultPage = n
	} else {
		defaultPage = n - 1
	}
	pages = defaultPage
	turnedPagesFromBack = 0
	shouldIBreakTheLoop := ""
	for pages >= 0 {
		if p == pages || p == pages+1 {
			shouldIBreakTheLoop = "yes"
		}
		if shouldIBreakTheLoop == "yes" {
			break
		}
		turnedPagesFromBack += 1
		pages -= 2
	}
	return turnedPagesFromBack
}

func PageCount(n, p int32) int32 {
	validity = CheckDataValidity(n, p)
	if validity == true {
		turnedPagesFromFront = FindNoOfTurnedPagesFromTheFront(n, p)
		turnedPagesFromBack = FindNoOfTurnedPagesFromTheBack(n, p)
		if turnedPagesFromFront < turnedPagesFromBack {
			minimumNoOfPagesTurned = turnedPagesFromFront
		} else {
			minimumNoOfPagesTurned = turnedPagesFromBack
		}
	}
	return minimumNoOfPagesTurned
}

func main() {
	n = 5
	p = 4
	minimumNoOfPagesTurned = PageCount(n, p)
	fmt.Println(minimumNoOfPagesTurned)
}
