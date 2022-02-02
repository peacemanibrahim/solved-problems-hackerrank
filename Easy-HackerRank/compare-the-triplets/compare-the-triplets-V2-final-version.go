// Link to problem statement: https://www.hackerrank.com/challenges/compare-the-triplets/problem
package main

import (
	"fmt"
)

func CompareTheTriplets(a, b [3]int32) [2]int32 {
	scores := [2]int32{0, 0}
	if a[0] >= 1 && a[0] <= 100 && b[0] >= 1 && b[0] <= 100 {
		fmt.Println("Valid data for first category")
		if a[0] > b[0] {
			fmt.Println("Alice is awarded one point for first category")
			aliceScore := scores[0] + 1
			scores[0] = aliceScore
		} else if a[0] < b[0] {
			fmt.Println("Bob is awarded one point for first category")
			bobScore := scores[1] + 1
			scores[1] = bobScore
		} else if a[0] == b[0] {
			fmt.Println("no points awarded to alice and bob for first category")
		}
	} else {
		fmt.Println("Invalid data for first category")
	}

	if a[1] >= 1 && a[1] <= 100 && b[1] >= 1 && b[1] <= 100 {
		fmt.Println("Valid data for second category")
		if a[1] > b[1] {
			fmt.Println("Alice is awarded one point for 2nd category")
			aliceScore := scores[0] + 1
			scores[0] = aliceScore
		} else if a[1] < b[1] {
			fmt.Println("Bob is awarded one point for 2nd category")
			bobScore := scores[1] + 1
			scores[1] = bobScore
		} else if a[1] == b[1] {
			fmt.Println("no points awarded to alice and bob for 2nd category")
		}
	} else {
		fmt.Println("Invalid data for second category")
	}

	if a[2] >= 1 && a[2] <= 100 && b[2] >= 1 && b[2] <= 100 {
		fmt.Println("Valid data for 3rd category")
		if a[2] > b[2] {
			fmt.Println("Alice is awarded one point for 3rd category")
			aliceScore := scores[0] + 1
			scores[0] = aliceScore
		} else if a[2] < b[2] {
			fmt.Println("Bob is awarded one point for 3rd category")
			bobScore := scores[1] + 1
			scores[1] = bobScore
		} else if a[2] == b[2] {
			fmt.Println("no points awarded to alice and bob for 3rd category")
		}
	} else {
		fmt.Println("Invalid data for 3rd category")
	}

	return scores
}

func main() {

	result := CompareTheTriplets([3]int32{1, 2, 3}, [3]int32{3, 2, 1})
	fmt.Println(result)

}
