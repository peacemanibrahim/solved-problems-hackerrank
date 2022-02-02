// Link to problem statement: https://www.hackerrank.com/challenges/kangaroo/problem
package main

import (
	"fmt"
	"log"
)

var (
	x1       int32
	v1       int32
	x2       int32
	v2       int32
	validity bool
	answer   string
)

func CheckDataValidity(x1 int32, v1 int32, x2 int32, v2 int32) bool {
	if x1 >= 0 && x2 > x1 && x2 <= 10000 && v1 >= 1 && v1 <= 10000 && v2 >= 1 && v2 <= 10000 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	validity = CheckDataValidity(x1, v1, x2, v2)
	if validity == true {
		firstKangarooPosition := x1 + v1
		secondKangarooPositon := x2 + v2
		for firstKangarooPosition <= 100000000 {
			if firstKangarooPosition == secondKangarooPositon {
				answer = "YES"
			}
			firstKangarooPosition = firstKangarooPosition + v1
			secondKangarooPositon = secondKangarooPositon + v2
		}
		if answer != "YES" {
			answer = "NO"
		}
	}
	return answer
}

func main() {
	x1 = 0
	v1 = 3
	x2 = 4
	v2 = 2

	result := Kangaroo(x1, v1, x2, v2)
	fmt.Println(result)
}

// Definition of terms:
// x1: First Kangaroo starting position
// x2: Second Kangaroo starting position
// v1: Distance covered in meters by first kangaroo per jump
// v2: Distance covered in meters by second kangaroo per jump
