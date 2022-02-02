package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

var (
	positiveCount    float64
	negativeCount    float64
	zeroCount        float64
	positiveFraction float64
	negativeFraction float64
	zeroFraction     float64
)

func plusMinus(arr []int32) {
	arrayLength := len(arr)
	positiveCount = 0
	negativeCount = 0
	zeroCount = 0

	if arrayLength > 0 && arrayLength < 100 {
		fmt.Println("Valid Array length")
		for i := range arr {
			if arr[i] >= -100 && arr[i] <= 100 {
				fmt.Println("Valid array data: ", arr[i])
				checkSign := math.Signbit(float64(arr[i]))
				if checkSign == true {
					fmt.Println("integer is a negative number: ", arr[i])
					negativeCount++
				} else if arr[i] == 0 {
					fmt.Println("integer is a zero value: ", arr[i])
					zeroCount++
				} else {
					fmt.Println("integer is a positive number: ", arr[i])
					positiveCount++
				}
			} else {
				log.Fatal("Invalid array data: ", arr[i])
			}

		}
	} else {
		log.Fatal("Invalid Array length: ", arrayLength)
	}

	positiveFraction = positiveCount / float64(arrayLength)
	fmt.Println("Positive integer fraction: ", strconv.FormatFloat(positiveFraction, 'f', 6, 64))

	negativeFraction = negativeCount / float64(arrayLength)
	fmt.Println("Negative integer fraction: ", strconv.FormatFloat(negativeFraction, 'f', 6, 64))

	zeroFraction = zeroCount / float64(arrayLength)
	fmt.Println("Zero integer fraction: ", strconv.FormatFloat(zeroFraction, 'f', 6, 64))

	return

}

func main() {
	arr1 := []int32{7, 1, 8, -3, -4, 0}

	plusMinus(arr1)

}
