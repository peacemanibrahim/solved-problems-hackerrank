// Link to problem statement: https://www.hackerrank.com/challenges/cut-the-sticks/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity                                bool
	newArray                                []int32
	listOfNoOfSticksLeftBeforeEachIteration []int32
)

func CheckDataValidity(arr []int32) bool {
	arrayLength := len(arr)
	if arrayLength >= 1 && arrayLength <= 1000 {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= 1000 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid length of array. Check the data.")
	}
	return validity
}

func Sort(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func CutTheSticks(arr []int32) []int32 {
	validity = CheckDataValidity(arr)
	if validity == true {
		Sort(arr)
		listOfNoOfSticksLeftBeforeEachIteration = append(listOfNoOfSticksLeftBeforeEachIteration, int32(len(arr)))
		for len(arr) >= 1 {
			for i := range arr {
				lenghtOfShortestStick := arr[0]
				resultOfCuttingShortestStickFromLongerSticks := arr[i] - lenghtOfShortestStick
				if resultOfCuttingShortestStickFromLongerSticks != 0 {
					newArray = append(newArray, resultOfCuttingShortestStickFromLongerSticks)
				}
			}
			lengthOfNewArray := int32(len(newArray))
			if lengthOfNewArray > 0 {
				listOfNoOfSticksLeftBeforeEachIteration = append(listOfNoOfSticksLeftBeforeEachIteration, lengthOfNewArray)
			}
			arr = newArray
			newArray = []int32{}
		}
	}
	return listOfNoOfSticksLeftBeforeEachIteration
}

func main() {
	arr := []int32{1, 2, 3, 4, 3, 3, 2, 1}

	listOfNoOfSticksLeftBeforeEachIteration = CutTheSticks(arr)
	fmt.Println(listOfNoOfSticksLeftBeforeEachIteration)
}
