// Link to problem statement: https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem
package main

import (
	"fmt"
	"log"
)

var (
	k                    int32
	remainingEnergyLevel int32
	validity             bool
)

func CheckDataValidity(c []int32, k int32) bool {
	arrayLengthOfC := int32(len(c))
	if arrayLengthOfC >= 2 && arrayLengthOfC <= 25 && k >= 1 && k <= arrayLengthOfC && arrayLengthOfC%k == 0 {
		for i := range c {
			if c[i] == 0 || c[i] == 1 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", c[i])
			}
		}
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func JumpingOnClouds(c []int32, k int32) int32 {
	remainingEnergyLevel = 100
	validity = CheckDataValidity(c, k)
	if validity == true {
		for i := 0; i < len(c); i += int(k) {
			if c[i] == 0 {
				remainingEnergyLevel -= 1
			} else if c[i] == 1 {
				remainingEnergyLevel -= 3
			}
		}
	}
	return remainingEnergyLevel
}

func main() {
	c := []int32{0, 0, 1, 0}
	k = 2

	remainingEnergyLevel = JumpingOnClouds(c, k)
	fmt.Println(remainingEnergyLevel)
}

// Note: c[n] is the cloud types along the path, whether cumulus represented as 0 or thunderhead represented as 1
// Note: k is the length of one jump
