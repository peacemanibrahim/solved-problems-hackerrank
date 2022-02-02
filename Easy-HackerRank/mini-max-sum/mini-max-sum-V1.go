package main

import (
	"fmt"
	"log"
	"math"
)

var validity bool

func checkArrayDataValidity(arr []int64) bool {
	for i := range arr {
		powerBase := int64(math.Pow(10, 9))
		if arr[i] >= 1 && arr[i] <= powerBase {
			fmt.Println("Valid array data:", arr[i])
			validity = true
		} else {
			log.Fatal("Invalid array data:", arr[i])
		}
	}
	return validity
}

func main() {
	arr := []int64{1, 2, 3, 4, 5}
	copyArr := arr
	var firstSumUp int64
	var secondSumUp int64
	var thirdSumUp int64
	var fourthSumUp int64
	firstSumUp = 0
	secondSumUp = 0
	thirdSumUp = 0
	fourthSumUp = 0

	validityResult := checkArrayDataValidity(arr)
	if validityResult == true {
		firstArray := append(arr[:0], arr[1:]...)
		for i := range firstArray {
			firstSumUp += firstArray[i]
		}

		fmt.Println(firstArray)
		fmt.Println(firstSumUp)

		// Reset the array. Let it contain all the numbers back
		arr = append(arr[:0], 1, 2, 3, 4, 5)
		fmt.Println(arr)

		secondArray := append(arr[:1], arr[2:]...)
		for i := range secondArray {
			secondSumUp += secondArray[i]
		}

		fmt.Println(secondArray)
		fmt.Println(secondSumUp)

		// Reset the array. Let it contain all the numbers back
		arr = append(arr[:0], 1, 2, 3, 4, 5)

		thirdArray := append(arr[:2], arr[3:]...)
		for i := range thirdArray {
			thirdSumUp += thirdArray[i]
		}

		fmt.Println(thirdArray)
		fmt.Println(thirdSumUp)

		// Reset the array. Let it contain all the numbers back
		arr = append(arr[:0], 1, 2, 3, 4, 5)

		fourthArray := append(arr[:3], arr[4:]...)

		for i := range fourthArray {
			fourthSumUp += fourthArray[i]
		}

		fmt.Println(fourthArray)
		fmt.Println(fourthSumUp)

		// Reset the array. Let it contain all the numbers back
		// arr = append(arr[:0])
		arr = append(arr[:0], copyArr...)
		fmt.Println(arr)

	} else {
		log.Fatal("Invalid Array Data")
	}

}
