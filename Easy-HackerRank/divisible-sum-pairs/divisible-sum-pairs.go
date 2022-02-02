// Link to problem statement: https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                        int32
	k                        int32
	validity                 bool
	countOfDivisibleSumPairs int32
)

func CheckDataValidity(n int32, k int32, ar []int32) bool {
	if n >= 2 && n <= 100 && k >= 1 && k <= 100 {
		for i := range ar {
			if ar[i] >= 1 && ar[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", ar[i])
			}
		}
	} else {
		log.Fatal("Invalid Data. Check the data")
	}
	return validity
}

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	validity = CheckDataValidity(n, k, ar)
	countOfDivisibleSumPairs = 0
	if validity == true {
		for len(ar) >= 2 {
			for i := 1; i < len(ar); i++ {
				sum := ar[0] + ar[i]
				fmt.Println(sum)
				sumDividedByK := sum % k
				if sumDividedByK == 0 {
					countOfDivisibleSumPairs++
				}
			}
			ar = ar[1:]
		}
	}
	return countOfDivisibleSumPairs
}

func main() {
	ar := []int32{1, 3, 2, 6, 1, 2}
	n = int32(len(ar))
	k = 3

	countOfDivisibleSumPairs = DivisibleSumPairs(n, k, ar)
	fmt.Println(countOfDivisibleSumPairs)
}
