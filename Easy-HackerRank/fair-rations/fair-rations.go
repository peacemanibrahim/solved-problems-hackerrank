// Link to problem statement: https://www.hackerrank.com/challenges/fair-rations/problem
package main

import (
	"fmt"
	"log"
	"strconv"
)

var (
	validity                      bool
	minimumNoOfLoavesToDistribute int32
)

func CheckDataValidity(B []int32) bool {
	arrayLengthOfB := len(B)
	if arrayLengthOfB >= 2 && arrayLengthOfB <= 1000 {
		for i := range B {
			if B[i] >= 1 && B[i] <= 10 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", B[i])
			}
		}
	} else {
		log.Fatal("Invalid array B length. Check the data.")
	}
	return validity
}

func FairRations(B []int32) string {
	validity = CheckDataValidity(B)
	if validity == true {
		for i := 0; i < len(B)-1; i++ {
			if B[i]%2 != 0 {
				B[i] = B[i] + 1
				B[i+1] = B[i+1] + 1
				minimumNoOfLoavesToDistribute += 2
			}
		}
		lastValueOfArrayB := B[len(B)-1]
		if lastValueOfArrayB%2 != 0 {
			return "NO"
		}
	}
	strRepOfMinimumNoOfLoavesToDistribute := strconv.Itoa(int(minimumNoOfLoavesToDistribute))
	return strRepOfMinimumNoOfLoavesToDistribute
}

func main() {
	B := []int32{2, 3, 4, 5, 6}

	result := FairRations(B)
	fmt.Println(result)
}
