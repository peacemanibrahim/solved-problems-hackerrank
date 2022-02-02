package main

import (
	"fmt"
	"log"
	"math"
)

var (
	n                    int32
	p                    int32
	validity             bool
	pages                int32
	turnedPagesFromFront int32
	turnedPagesFromBack  int32
	defaultPage          int32
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

func main() {
	n = 5
	p = 4
	validity = CheckDataValidity(n, p)

	if validity == true {
		defaultPage = 1
		pages = defaultPage
		turnedPagesFromFront = 0
		shouldIBreakTheLoop := ""
		for pages <= n+1 {
			if p == pages || p == pages-1 {
				fmt.Printf("Turning pages from the front: %d\n", turnedPagesFromFront)
				shouldIBreakTheLoop = "yes"
			}
			if shouldIBreakTheLoop == "yes" {
				break
			}
			turnedPagesFromFront += 1
			pages += 2
		}

		number := CheckIfNumberIsEvenOrOdd(n)
		if number == "even number" {
			defaultPage = n
		} else {
			defaultPage = n - 1
		}
		pages = defaultPage
		turnedPagesFromBack = 0
		shouldIBreakTheLoop = ""
		for pages >= 0 {
			if p == pages || p == pages+1 {
				fmt.Printf("Turning pages from the back: %d\n", turnedPagesFromBack)
				shouldIBreakTheLoop = "yes"
			}
			if shouldIBreakTheLoop == "yes" {
				break
			}
			turnedPagesFromBack += 1
			pages -= 2
		}
		if turnedPagesFromFront < turnedPagesFromBack {
			fmt.Printf("The minimum number of pages to turn is: %d\n", turnedPagesFromFront)
		} else {
			fmt.Printf("The minimum number of pages to turn is: %d\n", turnedPagesFromBack)
		}
	}
}

// n: This is the total number of pages in the book
// p: This is the page to be found
// Note: the n+1 on line 36 - the +1 after the n accounts for if there is only one page left at the back of the book
