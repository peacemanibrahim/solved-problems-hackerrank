package main

import (
	"fmt"
	"os"
)

var (
	sum_of_left_diagonal     int32
	sum_of_right_diagonal    int32
	difference_btw_diagonals int32
	validity                 bool
)

func sumUpLeftDiagonals(arr [][]int32) int32 {
	sum_of_left_diagonal = 0
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i][i])
		sum_of_left_diagonal += arr[i][i]
	}
	return sum_of_left_diagonal
}

func sumUpRightDiagonals(arr [][]int32) int32 {
	sum_of_right_diagonal = 0
	p := 0
	for j := 1; j <= len(arr); j++ {
		d := len(arr) - j
		fmt.Println(arr[p][d])
		sum_of_right_diagonal += arr[p][d]
		p += 1
	}
	return sum_of_right_diagonal
}

func checkValidityOfArrayData(arr [][]int32) bool {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr); y++ {
			if arr[x][y] >= -100 && arr[x][y] <= 100 && len(arr) == len(arr[x]) {
				fmt.Println("Valid data for arr: ", arr[x][y])
				validity = true
			} else {
				fmt.Println("Invalid data or length of array. Check your array data: ", arr[x][y])
				os.Exit(1)
			}
		}
	}
	return validity
}

func findDiagonalDifference(arr [][]int32) int32 {
	// Check the validity of the Array data
	validityResult := checkValidityOfArrayData(arr)

	if validityResult == true {
		// Calculate the sum of the left diagonal
		sum_of_left_diagonal = sumUpLeftDiagonals(arr)

		// Calculate the sum of the right diagonal
		sum_of_right_diagonal = sumUpRightDiagonals(arr)

		// Calculate the difference between diagonals
		difference_btw_diagonals = sum_of_left_diagonal - sum_of_right_diagonal
	}
	return difference_btw_diagonals
}

func main() {
	arrQuest := [][]int32{
		{5, 7, 8, 2, 3},
		{2, -8, 6, 5, 7},
		{3, 9, 2, 9, 1},
		{3, 9, 2, 6, 6},
		{3, 9, 2, 7, 6},
	}

	result := findDiagonalDifference(arrQuest)
	fmt.Println(result)

}
