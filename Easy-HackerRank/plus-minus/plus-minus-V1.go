package main

import (
	"fmt"
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

func main() {
	arr := []int32{4, 3, -5, 2, 9, 0, -4}
	arrayLength := len(arr)
	positiveCount = 0
	negativeCount = 0
	zeroCount = 0

	for i := range arr {
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
	}

	positiveFraction = positiveCount / float64(arrayLength)

	fmt.Println(strconv.FormatFloat(positiveFraction, 'x', 6, 64))

	negativeFraction = negativeCount / float64(arrayLength)

	fmt.Println(strconv.FormatFloat(negativeFraction, 'f', 6, 64))

	zeroFraction = zeroCount / float64(arrayLength)

	fmt.Println(strconv.FormatFloat(zeroFraction, 'f', 6, 64))

}
