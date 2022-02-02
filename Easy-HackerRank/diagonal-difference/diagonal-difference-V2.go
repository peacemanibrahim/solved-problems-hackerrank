package main

import (
	"fmt"
	"os"
)

func main() {
	arr := [][]int32{
		{5, 7, 8, 2, 3},
		{2, -8, 6, 5, 7},
		{3, 9, 2, 9, 1},
		{3, 9, 2, 6, 6},
		{3, 9, 2, 7, 6},
	}

	var (
		sum_of_left_diagonal     int32
		sum_of_right_diagonal    int32
		difference_btw_diagonals int32
		validity                 bool
	)

	// Check the validity of the Array data
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

	if validity == true {
		// Calculate the sum of the left diagonal
		sum_of_left_diagonal = 0
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i][i])
			sum_of_left_diagonal += arr[i][i]
		}

		// Calculate the sum of the right diagonal
		sum_of_right_diagonal = 0
		p := 0
		for j := 1; j <= len(arr); j++ {
			d := len(arr) - j
			fmt.Println(arr[p][d])
			sum_of_right_diagonal += arr[p][d]
			p += 1
		}

		// Calculate the difference between didagonals
		difference_btw_diagonals = sum_of_left_diagonal - sum_of_right_diagonal
		fmt.Println(difference_btw_diagonals)

	}
}
