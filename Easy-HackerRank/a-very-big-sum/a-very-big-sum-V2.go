package main

import (
	"fmt"
	"math"
)

func aVeryBigSum(no_of_elem int64, ar []int64) int64 {
	var sum int64
	sum = 0
	powerBase := int64(math.Pow(10, 10))
	if no_of_elem >= 1 && no_of_elem <= 10 {
		fmt.Println("Valid data for integer no_of_elem")
		for i := 0; i < len(ar); i++ {
			if ar[i] >= 0 && ar[i] <= powerBase {
				fmt.Println("Valid data for array ar: ", ar[i])
				sum = sum + ar[i]
			} else {
				fmt.Println("Invalid data for array ar: ", ar[i])
			}
		}
	} else {
		fmt.Println("Invalid data for integer no_of_elem")
	}

	return sum
}

func main() {

	result := aVeryBigSum(5, []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(result)

}
