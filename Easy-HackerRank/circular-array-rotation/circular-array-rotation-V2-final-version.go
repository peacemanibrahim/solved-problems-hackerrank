// Link to problem statement: https://www.hackerrank.com/challenges/circular-array-rotation/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	k                     int32
	validity              bool
	queriesResultList     []int32
	valueOfElementAtIndex int32
)

func CheckDataValidity(a []int32, k int32, queries []int32) bool {
	arrayLengthOfA := len(a)
	arrayLengthOfQueries := len(queries)
	basePower5 := int32(math.Pow(10, 5))
	if arrayLengthOfA >= 1 && arrayLengthOfA <= int(basePower5) && arrayLengthOfQueries >= 1 &&
		arrayLengthOfQueries <= 500 && k >= 1 && k <= basePower5 {
		for i := range a {
			if a[i] >= 1 && a[i] <= basePower5 {
				validity = true
			} else {
				log.Fatalf("Invalid a array data: %d\n", a[i])
			}
		}
		for i := range queries {
			if queries[i] >= 0 && queries[i] < int32(arrayLengthOfA) {
				validity = true
			} else {
				log.Fatalf("Invalid queries array data: %d\n", queries[i])
			}
		}
	} else {
		log.Fatal("Invalid Data. Check the data.")
	}
	return validity
}

func CircularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	arrayLengthOfQueries := len(queries)
	queriesResultList := make([]int32, arrayLengthOfQueries)
	validity = CheckDataValidity(a, k, queries)
	if validity == true {
		// The below "if k > arrayLenghtOfA { }" is because everytime k is equals to array length of |a|,
		// the elements of |a| come back to their original positions or indexes
		arrayLengthOfA := int32(len(a))
		if k > arrayLengthOfA {
			remainderOfKDividedByLengthOfA := k % arrayLengthOfA
			if remainderOfKDividedByLengthOfA == 0 {
				k = arrayLengthOfA
			} else {
				k = remainderOfKDividedByLengthOfA
			}
		}
		for i := 0; i < len(queries); i++ {
			indexOfReplacingElementCalculation := queries[i] - k
			isNegativeNumber := math.Signbit(float64(indexOfReplacingElementCalculation))
			if isNegativeNumber == true {
				indexOfReplacingElementAfterKRotations := len(a) + int(indexOfReplacingElementCalculation)
				valueOfElementAtIndex = a[indexOfReplacingElementAfterKRotations]
				queriesResultList[i] = valueOfElementAtIndex
			} else {
				indexOfReplacingElementAfterKRotations := indexOfReplacingElementCalculation
				valueOfElementAtIndex = a[indexOfReplacingElementAfterKRotations]
				queriesResultList[i] = valueOfElementAtIndex
			}
		}
	}
	return queriesResultList
}

func main() {
	a := []int32{1, 2, 3}
	k = 4
	queries := []int32{0, 1, 2}

	result := CircularArrayRotation(a, k, queries)
	fmt.Println(result)
}
