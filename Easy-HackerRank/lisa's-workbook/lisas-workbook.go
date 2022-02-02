// Link to problem statement: https://www.hackerrank.com/challenges/lisa-workbook/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                   int32
	k                   int32
	validity            bool
	pageNumber          int32
	specialProblemCount int32
)

func CheckDataValidity(n, k int32, arr []int32) bool {
	if n >= 1 && n <= 100 && k >= 1 && k <= 100 {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid data for n and/or k. Check the data.")
	}
	return validity
}

func Workbook(n, k int32, arr []int32) int32 {
	pageNumber = 1
	specialProblemCount = 0
	validity = CheckDataValidity(n, k, arr)
	if validity == true {
		for i := 0; i < len(arr); i++ {
			numberOfProblemsPerChapter := int(arr[i])
			for j := 1; j <= numberOfProblemsPerChapter; j++ {
				if j == int(pageNumber) {
					specialProblemCount++
				}
				if j%int(k) == 0 || j == numberOfProblemsPerChapter {
					pageNumber++
				}
				if j == numberOfProblemsPerChapter {
					numberOfProblemsPerChapter = 0
					break
				}

			}
		}
	}
	return specialProblemCount
}

func main() {
	n = 5
	k = 3
	arr := []int32{4, 2, 6, 1, 10}

	specialProblemCount = Workbook(n, k, arr)
	fmt.Println(specialProblemCount)
}
