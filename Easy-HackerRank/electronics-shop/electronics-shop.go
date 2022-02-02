// Link to problem statement: https://www.hackerrank.com/challenges/electronics-shop/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	validity          bool
	b                 int32
	mostExpensiveCost int32
)

func CheckDataValidity(keyboards, drives []int32, b int32) bool {
	basePower6 := int32(math.Pow(10, 6))
	keyboardsLength := len(keyboards)
	drivesLength := len(drives)
	if keyboardsLength >= 1 && keyboardsLength <= 1000 && drivesLength >= 1 && drivesLength <= 1000 && b >= 1 &&
		b <= basePower6 {
		for i := range keyboards {
			if keyboards[i] >= 1 && keyboards[i] <= basePower6 {
				validity = true
			} else {
				log.Fatalf("Invalid keyboard array data: %d\n", keyboards[i])
			}
		}
		for i := range drives {
			if drives[i] >= 1 && drives[i] <= basePower6 {
				validity = true
			} else {
				log.Fatalf("Invalid drives array data: %d\n", drives[i])
			}
		}
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func GetMoneySpent(keyboards, drives []int32, b int32) int32 {
	validity = CheckDataValidity(keyboards, drives, b)
	if validity == true {
		mostExpensiveCost = -1
		for i := 0; i < len(keyboards); i++ {
			for j := 0; j < len(drives); j++ {
				cost := keyboards[i] + drives[j]
				if cost > mostExpensiveCost && cost <= b {
					mostExpensiveCost = cost
				}
			}
		}
	}
	return mostExpensiveCost
}

func main() {
	keyboards := []int32{40, 50, 60}
	drives := []int32{5, 8, 12}
	b = 60

	mostExpensiveCost = GetMoneySpent(keyboards, drives, b)
	fmt.Println(mostExpensiveCost)
}

// Note: b means budget
