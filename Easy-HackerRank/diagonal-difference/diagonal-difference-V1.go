package main

import (
	"fmt"
	"log"
)

func main() {
	arr := [][]int32{
		{5, 7, 8, 2, 3},
		{2, -3, 6, 5, 7},
		{3, 9, 2, 7, 9},
		{3, 9, 2, 7, 6},
		{3, 9, 2, 7, 6},
	}

	var (
		sum_of_left_diagonal     int32
		sum_of_right_diagonal    int32
		difference_btw_diagonals int32
		validity                 bool
	)

	// Check the validity of the Array
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr); y++ {
			if arr[x][y] >= -100 && arr[x][y] <= 100 {
				fmt.Println("Valid data for arr: ", arr[x][y])
				validity = true
			} else {
				log.Fatal("Invalid data for array. Check your array data")
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
	} else {
		fmt.Println("Invalid data for arr")
	}
}
