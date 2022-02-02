// Link to problem statement: https://www.hackerrank.com/challenges/simple-array-sum/problem
package main

import (
	"fmt"
	"log"
)

var (
	sumOfArray int32
)

func SimpleArraySum(arr []int32) int32 {
	arrayLength := len(arr)

	if arrayLength > 0 && arrayLength <= 1000 {
		for i := 0; i < arrayLength; i++ {
			if arr[i] > 0 && arr[i] <= 1000 {
				sumOfArray += arr[i]
			} else {
				log.Fatalf("Invalid Array Data %d\n", arr[i])
			}
		}
	} else {
		log.Fatal("Invalid Array Length")
	}
	return sumOfArray
}

func main() {

	arr := []int32{23, 54, 1000, 45}

	sumOfArray = SimpleArraySum(arr)
	fmt.Println(sumOfArray)

}
