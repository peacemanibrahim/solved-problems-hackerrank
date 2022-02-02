// Link to problem statement: https://www.hackerrank.com/challenges/strange-advertising/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                                 int32
	validity                          bool
	currentDayLikes                   int32
	totalLikesOfAdByTheEndOfGivenDays int32
	followingDayRecipients            int32
)

func CheckDataValidity(n int32) bool {
	if n >= 1 && n <= 50 {
		validity = true
	} else {
		log.Fatalf("Invalid data for n: %d\n", n)
	}
	return validity
}

func ViralAdvertising(n int32) int32 {
	if n == 1 {
		return 2
	}
	validity = CheckDataValidity(n)
	if validity == true {
		totalLikesOfAdByTheEndOfGivenDays = 2
		followingDayRecipients = 6
		for i := 2; i <= int(n); i++ {
			currentDayLikes = followingDayRecipients / 2
			totalLikesOfAdByTheEndOfGivenDays += currentDayLikes
			followingDayRecipients = currentDayLikes * 3
			currentDayLikes = 0
		}
	}
	return totalLikesOfAdByTheEndOfGivenDays
}

func main() {
	n = 1
	totalLikesOfAdByTheEndOfGivenDays = ViralAdvertising(n)
	fmt.Println(totalLikesOfAdByTheEndOfGivenDays)
}
