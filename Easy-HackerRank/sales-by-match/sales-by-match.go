// Link to problem statement: https://www.hackerrank.com/challenges/sock-merchant/problem
package main

import (
	"fmt"
	"log"
)

var (
	n                                     int32
	validity                              bool
	socksOfSameColorCount                 int32
	currentNoOfSameColorSocksPairs        int32
	totalNoOfSocksPairsWithMatchingColors int32
)

func CheckDataValidity(n int32, ar []int32) bool {
	if n >= 1 && n <= 100 {
		for i := range ar {
			if ar[i] >= 1 && ar[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", ar[i])
			}
		}
	} else {
		log.Fatal("Invalid array length. Check the array length.")
	}
	return validity
}

func SockMerchant(n int32, ar []int32) int32 {
	validity = CheckDataValidity(n, ar)
	if validity == true {
		totalNoOfSocksPairsWithMatchingColors = 0
		for i := 1; i <= 100; i++ {
			socksOfSameColorCount = 0
			currentNoOfSameColorSocksPairs = 0
			for j := range ar {
				if ar[j] == int32(i) {
					socksOfSameColorCount++
				}
			}
			currentNoOfSameColorSocksPairs = socksOfSameColorCount / 2
			totalNoOfSocksPairsWithMatchingColors += currentNoOfSameColorSocksPairs
		}
	}
	return totalNoOfSocksPairsWithMatchingColors
}

func main() {
	ar := []int32{1, 2, 1, 2, 1, 3, 2}
	n = int32(len(ar))

	totalNoOfSocksPairsWithMatchingColors = SockMerchant(n, ar)
	fmt.Println(totalNoOfSocksPairsWithMatchingColors)
}
