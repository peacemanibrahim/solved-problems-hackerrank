// Link to the problem statement: https://www.hackerrank.com/challenges/equality-in-a-array/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity                    bool
	numberOccurrenceCount       int32
	numberOccurringTheMost      int32
	minimumNoOfElementsToDelete int32
)

func CheckDataValidity(arr []int32) bool {
	arrayLength := len(arr)
	if arrayLength >= 1 && arrayLength <= 100 {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid array length. Check the data.")
	}
	return validity
}

func EqualizeArray(arr []int32) int32 {
	arrayLength := int32(len(arr))
	validity = CheckDataValidity(arr)
	if validity == true {
		numberOccurrenceCount = 0
		numberOccurringTheMost = 0
		for i := 1; i <= 100; i++ {
			for j := 0; j < len(arr); j++ {
				if i == int(arr[j]) {
					numberOccurrenceCount++
				}
			}
			if numberOccurrenceCount > numberOccurringTheMost {
				numberOccurringTheMost = numberOccurrenceCount
			}
			numberOccurrenceCount = 0
		}
		minimumNoOfElementsToDelete = arrayLength - numberOccurringTheMost
	}
	return minimumNoOfElementsToDelete
}

func main() {
	arr := []int32{3, 3, 2, 1, 3}

	minimumNoOfElementsToDelete = EqualizeArray(arr)
	fmt.Println(minimumNoOfElementsToDelete)
}
