package main

import (
	"fmt"
	"log"
	"math"
)

var (
	validity    bool
	firstSumUp  int32
	secondSumUp int32
	thirdSumUp  int32
	fourthSumUp int32
	fifthSumUp  int32
)

func checkArrayDataValidity(arr []int32) bool {
	for i := range arr {
		powerBase := int32(math.Pow(10, 9))
		if arr[i] >= 1 && arr[i] <= powerBase {
			fmt.Println("Valid array data:", arr[i])
			validity = true
		} else {
			log.Fatal("Invalid array data:", arr[i])
		}
	}
	return validity
}

func findMinAndMax(sumUpArray []int32) (min int32, max int32) {
	min = sumUpArray[0]
	max = sumUpArray[0]

	for _, value := range sumUpArray {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}
	return min, max
}

func miniMaxSum(arr []int32) {
	sumUpArray := []int32{}
	firstSumUp = 0
	secondSumUp = 0
	thirdSumUp = 0
	fourthSumUp = 0
	fifthSumUp = 0

	validityResult := checkArrayDataValidity(arr)
	if validityResult == true {
		firstSumUp = arr[1] + arr[2] + arr[3] + arr[4]
		secondSumUp = arr[0] + arr[2] + arr[3] + arr[4]
		thirdSumUp = arr[0] + arr[1] + arr[3] + arr[4]
		fourthSumUp = arr[0] + arr[1] + arr[2] + arr[4]
		fifthSumUp = arr[0] + arr[1] + arr[2] + arr[3]

		sumUpArray = append(sumUpArray, firstSumUp, secondSumUp, thirdSumUp, fourthSumUp, fifthSumUp)
		min, max := findMinAndMax(sumUpArray)
		fmt.Println(min, max)
	} else {
		log.Fatal("Invalid Array Data")
	}
	return
}

func main() {
	arr1 := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr1)
}
