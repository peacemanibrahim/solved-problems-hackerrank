package main

import (
	"fmt"
	"log"
	"math"
)

var validity bool

func checkArrayDataValidity(arr [5]int64) bool {
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

func calculateMinAndMax(sumUpArray []int64) (min int64, max int64) {
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

func main() {
	arr := [5]int64{1, 2, 3, 4, 5}
	sumUpArray := []int64{}
	var firstSumUp int64
	var secondSumUp int64
	var thirdSumUp int64
	var fourthSumUp int64
	var fifthSumUp int64
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

		fmt.Println(firstSumUp)
		fmt.Println(secondSumUp)
		fmt.Println(thirdSumUp)
		fmt.Println(fourthSumUp)
		fmt.Println(fifthSumUp)

		fmt.Println(sumUpArray)

		min, max := calculateMinAndMax(sumUpArray)
		fmt.Println(min, max)

	} else {
		log.Fatal("Invalid Array Data")
	}

}
