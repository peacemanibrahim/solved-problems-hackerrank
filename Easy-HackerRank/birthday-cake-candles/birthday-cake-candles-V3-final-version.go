// Link to problem statement: https://www.hackerrank.com/challenges/birthday-cake-candles/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	validity        bool
	nieceAge        int32
	blownOutCandles int32
	tallestCandle   int32
)

func checkArrayValidity(arr []int32) bool {
	nieceAge = int32(len(arr))
	power5Base := int32(math.Pow(10, 5))
	power7Base := int32(math.Pow(10, 7))
	if nieceAge >= 1 && nieceAge <= power5Base {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= power7Base {
				validity = true
			} else {
				log.Fatal("Invalid array data:", arr[i])
			}
		}
	}
	return validity
}

func countTallestCandlesOccurence(arr []int32) int32 {
	var count int32
	count = 0
	for i := range arr {
		if arr[i] == tallestCandle {
			count++
		}
	}
	return count
}

func birthdayCakeCandles(arr []int32) int32 {
	validityResult := checkArrayValidity(arr)
	if validityResult == true {
		tallestCandle = arr[0]
		for i := range arr {
			if arr[i] > tallestCandle {
				tallestCandle = arr[i]
			}
		}
		blownOutCandles = countTallestCandlesOccurence(arr)
	} else {
		log.Fatal("Invalid Array Data")
	}

	return blownOutCandles
}

func main() {
	arr := []int32{4, 3, 4, 2, 5, 6, 6}
	result := birthdayCakeCandles(arr)
	fmt.Println(result)
}
