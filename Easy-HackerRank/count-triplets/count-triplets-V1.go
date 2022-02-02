// Link to problem statement on Hackerrank: https://www.hackerrank.com/challenges/count-triplets-1/problem
// All test cases did Not pass on hackerrank.com
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	r                                           int64
	validity                                    bool
	secondValueOfCurrentTriplet                 int64
	thirdValueOfCurrentTriplet                  int64
	noOfOccurrenceOfSecondValueOfCurrentTriplet int64
	noOfOccurrenceOfThirdValueOfCurrentTriplet  int64
	noOfTimesTheCurrentTripletOccurs            int64
	tripletsSatisfyingCriteriaCount             int64
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
		mapOfArr := make(map[int64]int64)
		for _, v := range arr {
			mapOfArr[v]++
		}
		for len(arr) > 2 {
			firstValueOfCurrentTriplet := arr[0]
			secondValueOfCurrentTriplet = firstValueOfCurrentTriplet * r
			thirdValueOfCurrentTriplet = secondValueOfCurrentTriplet * r
			// delete the first element of the array in mapOfArr
			noOfOccurrenceOfFirstElementOfArrayInTheMap := mapOfArr[arr[0]]
			if noOfOccurrenceOfFirstElementOfArrayInTheMap > 1 {
				mapOfArr[arr[0]] = noOfOccurrenceOfFirstElementOfArrayInTheMap - 1
			} else if noOfOccurrenceOfFirstElementOfArrayInTheMap == 1 {
				delete(mapOfArr, arr[0])
			}
			noOfOccurrenceOfSecondValueOfCurrentTriplet = mapOfArr[secondValueOfCurrentTriplet]
			noOfOccurrenceOfThirdValueOfCurrentTriplet = mapOfArr[thirdValueOfCurrentTriplet]
			if noOfOccurrenceOfSecondValueOfCurrentTriplet != 0 && noOfOccurrenceOfThirdValueOfCurrentTriplet != 0 {
				if secondValueOfCurrentTriplet != thirdValueOfCurrentTriplet {
					noOfTimesTheCurrentTripletOccurs = noOfOccurrenceOfSecondValueOfCurrentTriplet *
						noOfOccurrenceOfThirdValueOfCurrentTriplet
					tripletsSatisfyingCriteriaCount += noOfTimesTheCurrentTripletOccurs
				} else if secondValueOfCurrentTriplet == thirdValueOfCurrentTriplet &&
					noOfOccurrenceOfSecondValueOfCurrentTriplet > 1 {
					n := noOfOccurrenceOfSecondValueOfCurrentTriplet
					numerator := n * (n - 1)
					denominator := 2 * (2 - 1)
					nCombination2 := numerator / int64(denominator)
					noOfTimesTheCurrentTripletOccurs = nCombination2
					tripletsSatisfyingCriteriaCount += noOfTimesTheCurrentTripletOccurs
				}
			}
			noOfTimesTheCurrentTripletOccurs = 0
			arr = arr[1:]
		}
	}
	return tripletsSatisfyingCriteriaCount
}

func main() {
	arr := []int64{1, 2, 1, 2, 4}
	r = 2

	tripletsSatisfyingCriteriaCount = CountTriplets(arr, r)
	fmt.Println(tripletsSatisfyingCriteriaCount)
}
