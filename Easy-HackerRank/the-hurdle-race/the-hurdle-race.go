// Link to problem statement: https://www.hackerrank.com/challenges/the-hurdle-race/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity                  bool
	k                         int32
	noOfDosesOfPotionRequired int32
)

func CheckDataValidity(k int32, height []int32) bool {
	arrayLength := len(height)
	if arrayLength >= 1 && arrayLength <= 100 && k >= 1 && k <= 100 {
		for i := range height {
			if height[i] >= 1 && height[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", height[i])
			}
		}
	} else {
		log.Fatal("Invalid data for array length and/or k. Check the data.")
	}
	return validity
}

func HurdleRace(k int32, height []int32) int32 {
	validity = CheckDataValidity(k, height)
	if validity == true {
		maximumHeightOfHurdles := height[0]
		for i := 1; i < len(height); i++ {
			if height[i] > maximumHeightOfHurdles {
				maximumHeightOfHurdles = height[i]
			}
		}
		if maximumHeightOfHurdles > k {
			noOfDosesOfPotionRequired = maximumHeightOfHurdles - k
		} else if maximumHeightOfHurdles <= k {
			noOfDosesOfPotionRequired = 0
		}
	}
	return noOfDosesOfPotionRequired
}

func main() {
	height := []int32{2, 5, 4, 5, 2}
	k = 7

	noOfDosesOfPotionRequired = HurdleRace(k, height)
	fmt.Println(noOfDosesOfPotionRequired)
}
