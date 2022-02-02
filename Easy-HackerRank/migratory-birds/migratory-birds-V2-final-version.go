// Link to problem statement: https://www.hackerrank.com/challenges/migratory-birds/problem
package main

import (
	"fmt"
	"log"
)

var (
	oneCount                    int32
	twoCount                    int32
	threeCount                  int32
	fourCount                   int32
	fiveCount                   int32
	validity                    bool
	birdTypeIdCountList         []int32
	mostFrequentBirdTypeIdCount int32
	mostFrequentBirdTypeId      int32
)

const (
	one   = 1
	two   = 2
	three = 3
	four  = 4
	five  = 5
)

func CheckDataValidity(arr []int32) bool {
	arrayLength := len(arr)
	if arrayLength >= 5 && arrayLength <= 200000 {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= 5 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid array length")
	}
	return validity
}

func FindNoOfOccurenceOfBirdTypeIds(arr []int32) []int32 {
	oneCount = 0
	twoCount = 0
	threeCount = 0
	fourCount = 0
	fiveCount = 0
	birdTypeIdCountList := make([]int32, 5)
	for i := range arr {
		if arr[i] == one {
			oneCount++
		} else if arr[i] == two {
			twoCount++
		} else if arr[i] == three {
			threeCount++
		} else if arr[i] == four {
			fourCount++
		} else {
			fiveCount++
		}
	}
	birdTypeIdCountList[0] = oneCount
	birdTypeIdCountList[1] = twoCount
	birdTypeIdCountList[2] = threeCount
	birdTypeIdCountList[3] = fourCount
	birdTypeIdCountList[4] = fiveCount

	return birdTypeIdCountList
}

func MigratoryBirds(arr []int32) int32 {
	validity = CheckDataValidity(arr)
	if validity == true {
		birdTypeIdCountList = FindNoOfOccurenceOfBirdTypeIds(arr)
		mostFrequentBirdTypeIdCount = birdTypeIdCountList[0]
		birdTypeId := 1
		mostFrequentBirdTypeId = int32(birdTypeId)
		for i := 1; i < len(birdTypeIdCountList); i++ {
			birdTypeId++
			if mostFrequentBirdTypeIdCount < birdTypeIdCountList[i] {
				mostFrequentBirdTypeIdCount = birdTypeIdCountList[i]
				mostFrequentBirdTypeId = int32(birdTypeId)
			}
		}
	}
	return mostFrequentBirdTypeId
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	result := MigratoryBirds(arr)
	fmt.Println(result)
}
