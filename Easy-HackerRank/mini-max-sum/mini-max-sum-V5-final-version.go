// Link to problem statement: https://www.hackerrank.com/challenges/mini-max-sum/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	validity    bool
	firstSumUp  int64
	secondSumUp int64
	thirdSumUp  int64
	fourthSumUp int64
	fifthSumUp  int64
)

func checkArrayDataValidity(arr []int64) bool {
	for i := range arr {
		powerBase := int64(math.Pow(10, 9))
		if arr[i] >= 1 && arr[i] <= powerBase && len(arr) == 5 {
			validity = true
		} else {
			log.Fatal("Invalid array data:", arr[i])
		}
	}
	return validity
}

func findMinAndMax(sumUpArray []int64) (min int64, max int64) {
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

func miniMaxSum(arr []int64) {
	sumUpArray := []int64{}
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
	// Note: I used 64 bit integers to avoid integer overflow(which occurs if I use 32 bit integer)
	arr1 := []int64{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	miniMaxSum(arr1)
}
