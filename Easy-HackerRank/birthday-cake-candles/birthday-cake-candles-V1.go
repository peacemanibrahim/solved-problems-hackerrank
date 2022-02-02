package main

import (
	"fmt"
	"log"
	"math"
)

var validity bool
var nieceAge int32
var tallestCandle int32

func checkArrayValidity(arr []int32) bool {
	nieceAge = int32(len(arr))
	power5Base := int32(math.Pow(10, 5))
	power7Base := int32(math.Pow(10, 7))
	if nieceAge >= 1 && nieceAge <= power5Base {
		for i := range arr {
			if arr[i] >= 1 && arr[i] <= power7Base {
				fmt.Println("Valid array data:", arr[i])
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

func main() {
	arr := []int32{4, 3, 4, 2, 5, 6, 6}
	validityResult := checkArrayValidity(arr)

	if validityResult == true {
		tallestCandle = arr[0]
		for i := range arr {
			if arr[i] > tallestCandle {
				tallestCandle = arr[i]
			}
		}
		fmt.Println(tallestCandle)
		tallestCandleOccurence := countTallestCandlesOccurence(arr)
		fmt.Println(tallestCandleOccurence)
	} else {
		log.Fatal("Invalid Array Data")
	}

}
