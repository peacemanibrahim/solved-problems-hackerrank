// Link to problem statement on Hackerrank: https://www.hackerrank.com/challenges/count-triplets-1/problem
// This is the count triplets solution implemented using "maps" data structures and dynamic programming.
// All test cases passed on hackerrank.com
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	r                                          int64
	validity                                   bool
	firstValueOfCurrentTriplet                 int64
	secondValueOfCurrentTriplet                int64
	thirdValueOfCurrentTriplet                 int64
	noOfOccurrenceOfFirstValueOfCurrentTriplet int64
	noOfOccurrenceOfThirdValueOfCurrentTriplet int64
	noOfTimesTheCurrentTripletOccurs           int64
	tripletsSatisfyingCriteriaCount            int64
)

func CheckDataValidity(arr []int64, r int64) bool {
	arrayLength := len(arr)
	basePower9 := int64(math.Pow(10, 9))
	if arrayLength >= 1 && arrayLength <= 100000 && r >= 1 && r <= basePower9 {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= basePower9 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid length of array and/or value of r. Check the data.")
	}
	return validity
}

func CountTriplets(arr []int64, r int64) int64 {
	validity = CheckDataValidity(arr, r)
	if validity == true {
		leftMap := make(map[int64]int64)
		for _, v := range arr {
			leftMap[v] = 0
		}
		rightMap := make(map[int64]int64)
		for _, v := range arr {
			rightMap[v]++
		}

		for i := 0; i < len(arr); i++ {
			// delete arr[i] from right map
			rightMap[arr[i]]--
			secondValueOfCurrentTriplet = arr[i]
			if secondValueOfCurrentTriplet%r == 0 {
				firstValueOfCurrentTriplet = secondValueOfCurrentTriplet / r
			}
			thirdValueOfCurrentTriplet = secondValueOfCurrentTriplet * r
			noOfOccurrenceOfFirstValueOfCurrentTriplet = leftMap[firstValueOfCurrentTriplet]
			noOfOccurrenceOfThirdValueOfCurrentTriplet = rightMap[thirdValueOfCurrentTriplet]
			noOfTimesTheCurrentTripletOccurs = noOfOccurrenceOfFirstValueOfCurrentTriplet *
				noOfOccurrenceOfThirdValueOfCurrentTriplet
			tripletsSatisfyingCriteriaCount += noOfTimesTheCurrentTripletOccurs
			firstValueOfCurrentTriplet = 0
			// add arr[i] to left map
			leftMap[arr[i]]++
		}
	}
	return tripletsSatisfyingCriteriaCount
}

func main() {
	arr := []int64{1, 3, 3, 9, 3, 27, 81}
	r = 3

	tripletsSatisfyingCriteriaCount = CountTriplets(arr, r)
	fmt.Println(tripletsSatisfyingCriteriaCount)
}
