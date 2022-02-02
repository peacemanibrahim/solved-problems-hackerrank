// Link to problem statement: https://www.hackerrank.com/challenges/between-two-sets/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity               bool
	numberBetweenTwoArrays int32
	breakLoop              string
)

func CheckDataValidity(a, b []int32) bool {
	arrayLengthOfA := len(a)
	arrayLengthOfB := len(b)
	if arrayLengthOfA >= 1 && arrayLengthOfA <= 10 && arrayLengthOfB >= 1 && arrayLengthOfB <= 10 {
		for i := range a {
			if a[i] >= 1 && a[i] <= 100 {

			} else {
				log.Fatalf("Invalid data for array a, %d\n", a[i])
			}
		}

		for i := range b {
			if b[i] >= 1 && b[i] <= 100 {

			} else {
				log.Fatalf("Invalid data for array b, %d\n", b[i])
			}
		}
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func GetTotalNumberBetweenArrays(a []int32, b []int32) int32 {
	validity = CheckDataValidity(a, b)
	numberBetweenTwoArrays = 0
	if validity == true {
		for numberBeingConsidered := 1; numberBeingConsidered <= 100; numberBeingConsidered++ {
			breakLoop = "NO"
			for i := range a {
				numberBeingConsideredFactorsCheck := int32(numberBeingConsidered) % a[i]
				if numberBeingConsideredFactorsCheck != 0 {
					breakLoop = "YES"
					break
				}
			}
			if breakLoop == "YES" {
				continue
			}

			breakLoop = "NO"
			for i := range b {
				arraybElementFactorsCheck := b[i] % int32(numberBeingConsidered)
				if arraybElementFactorsCheck != 0 {
					breakLoop = "YES"
					break
				}
			}
			if breakLoop == "YES" {
				continue
			}
			numberBetweenTwoArrays++
		}
	}
	return numberBetweenTwoArrays
}

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	numberBetweenTwoArrays = GetTotalNumberBetweenArrays(a, b)
	fmt.Println(numberBetweenTwoArrays)
}
