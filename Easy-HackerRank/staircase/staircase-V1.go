package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var staircaseSize int32
	// var staircaseHeight int32
	// var staircaseWidth int32

	staircaseSize = 8
	// staircaseWidth := []string{}

	if staircaseSize > 0 && staircaseSize <= 100 {
		fmt.Println("Valid data for staircase size:", staircaseSize)
		for i := 1; i <= int(staircaseSize); i++ {
			// staircaseWidth = append(staircaseWidth, "#")
			// fmt.Printf(staircaseWidth)
			fmt.Printf("%8s\n", strings.Repeat("#", i))
		}
	} else {
		log.Fatal("Invalid data for staircase size: ", staircaseSize)
	}
}
