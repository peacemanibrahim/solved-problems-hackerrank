// Link to problem statement: https://www.hackerrank.com/challenges/service-lane/problem
package main

import (
	"fmt"
	"log"
)

var (
	n        int32
	validity bool
)

func CheckDataValidity(width []int32, cases [][]int32) bool {
	arrayLengthOfWidth := len(width)
	arrayLengthOfCases := len(cases)
	if arrayLengthOfWidth >= 2 && arrayLengthOfWidth <= 100000 && arrayLengthOfCases >= 1 &&
		arrayLengthOfCases <= 1000 {
		for k := range width {
			if width[k] >= 1 && width[k] <= 3 {
				validity = true
			} else {
				log.Fatalf("Invalid width array data: %d\n", width[k])
			}
		}
		for i := 0; i < len(cases); i++ {
			j := 0
			entryIndices := cases[i][j]
			exitIndices := cases[i][j+1]
			if entryIndices >= 0 && exitIndices > entryIndices && exitIndices < int32(arrayLengthOfWidth) &&
				entryIndices <= int32(arrayLengthOfWidth)-1 && exitIndices <= int32(arrayLengthOfWidth)-1 {
				validity = true
			} else {
				log.Fatalf("Invalid cases array data: %d\n", cases[i])
			}
		}
	} else {
		log.Fatal("Invalid array length for width and/or cases. Check the data.")
	}
	return validity
}

func ServiceLane(n int32, width []int32, cases [][]int32) []int32 {
	validity = CheckDataValidity(width, cases)
	listOfMaxWidthOfVehicleThatCanTravelSegment := []int32{}
	if validity == true {
		for i := 0; i < len(cases); i++ {
			j := 0
			entryIndices := cases[i][j]
			exitIndices := cases[i][j+1]
			minimumWidthOfEachSegment := width[entryIndices]
			for k := entryIndices; k <= exitIndices; k++ {
				if minimumWidthOfEachSegment > width[k] {
					minimumWidthOfEachSegment = width[k]
				}
			}
			listOfMaxWidthOfVehicleThatCanTravelSegment = append(listOfMaxWidthOfVehicleThatCanTravelSegment,
				minimumWidthOfEachSegment)
			minimumWidthOfEachSegment = 0
		}
	}
	return listOfMaxWidthOfVehicleThatCanTravelSegment
}

func main() {
	n = 8
	width := []int32{2, 3, 1, 2, 3, 2, 3, 3}
	cases := [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}

	result := ServiceLane(n, width, cases)
	fmt.Println(result)
}

// Note: n is the length of width.
