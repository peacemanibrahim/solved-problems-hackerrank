// Link to problem statement: https://www.hackerrank.com/challenges/a-very-big-sum/problem
package main

import (
	"fmt"
	"math"
)

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	sum = 0
	powerBase := int64(math.Pow(10, 10))
	for i := 0; i < len(ar); i++ {
		if ar[i] >= 0 && ar[i] <= powerBase {
			fmt.Println("Valid data for array ar: ", ar[i])
			sum = sum + ar[i]
		} else {
			fmt.Println("Invalid data for array ar: ", ar[i])
		}
	}

	return sum
}

func main() {

	result := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(result)

}
