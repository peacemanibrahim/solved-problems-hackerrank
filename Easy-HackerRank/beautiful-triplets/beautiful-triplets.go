// Link to problem statement: https://www.hackerrank.com/challenges/beautiful-triplets/problem
package main

import (
	"fmt"
	"log"
)

var (
	d                                 int32
	validity                          bool
	secondValueOfEachBeautifulTriplet int32
	beautifulTripletsCount            int32
)

func CheckDataValidity(d int32, arr []int32) bool {
	arrayLength := len(arr)
	if arrayLength >= 1 && arrayLength <= 10000 && d >= 1 && d <= 20 {
		for i := range arr {
			if arr[i] >= 0 && arr[i] <= 20000 {
				validity = true
			} else {
				log.Fatalf("Invalid arra data: %d\n", arr[i])
			}
		}
		for j := 1; j < len(arr); j++ {
			if arr[j] >= arr[j-1] {
				validity = true
			} else {
				log.Fatal("Invalid array sequence. Check the array.")
			}
		}
	} else {
		log.Fatal("Invalid data for array length and/or d. Check the data.")
	}
	return validity
}

func BeautifulTriplets(d int32, arr []int32) int32 {
	secondValueOfEachBeautifulTriplet = 0
	beautifulTripletsCount = 0
	validity = CheckDataValidity(d, arr)
	if validity == true {
		var isSecondValueOfBeautifulTripletGotten bool
		for len(arr) > 2 {
			for i := 1; i < len(arr); i++ {
				if arr[i]-arr[0] == d {
					secondValueOfEachBeautifulTriplet = arr[i]
					isSecondValueOfBeautifulTripletGotten = true
				}
				if arr[i]-secondValueOfEachBeautifulTriplet == d && isSecondValueOfBeautifulTripletGotten == true {
					beautifulTripletsCount++
					break
				}
			}
			secondValueOfEachBeautifulTriplet = 0
			isSecondValueOfBeautifulTripletGotten = false
			arr = arr[1:]
		}
	}
	return beautifulTripletsCount
}

func main() {
	arr := []int32{0, 1, 2, 4, 5, 11, 14, 15, 16, 17, 18, 19, 21, 23, 26, 27, 29, 31, 33, 34, 36, 37, 38}
	d = 5

	beautifulTripletsCount = BeautifulTriplets(d, arr)
	fmt.Println(beautifulTripletsCount)
}
