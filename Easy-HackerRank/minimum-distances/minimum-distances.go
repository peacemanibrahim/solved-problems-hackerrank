// Link to problem statement: https://www.hackerrank.com/challenges/minimum-distances/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity                       bool
	distanceBtwTwoEqualArrayValues int32
	finalMinimumDistance           int32
)

func CheckDataValidity(a []int32) bool {
	arrayLength := len(a)
	if arrayLength >= 1 && arrayLength <= 1000 {
		for i := range a {
			if a[i] >= 1 && a[i] <= 100000 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", a[i])
			}
		}
	} else {
		log.Fatal("Invalid array length. Check the data.")
	}
	return validity
}

func MinimumDistances(a []int32) int32 {
	finalMinimumDistance = 99999
	distanceBtwTwoEqualArrayValues = finalMinimumDistance
	validity = CheckDataValidity(a)
	if validity == true {
		for i := 0; i < len(a)-1; i++ {
			for j := i + 1; j < len(a); j++ {
				if a[i] == a[j] {
					distanceBtwTwoEqualArrayValues = int32(j - i)
					break
				}
			}
			if distanceBtwTwoEqualArrayValues < finalMinimumDistance {
				finalMinimumDistance = distanceBtwTwoEqualArrayValues
				distanceBtwTwoEqualArrayValues = 99999
			}
		}
		if finalMinimumDistance == 99999 {
			finalMinimumDistance = -1
		}
	}
	return finalMinimumDistance
}

func main() {
	a := []int32{3, 2, 1, 2, 3}

	finalMinimumDistance = MinimumDistances(a)
	fmt.Println(finalMinimumDistance)
}
