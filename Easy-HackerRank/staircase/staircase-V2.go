package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func createStaircase(staircaseSize int32) {
	if staircaseSize > 0 && staircaseSize <= 100 {
		fmt.Println("Valid data for staircase size:", staircaseSize)
		for i := 1; i <= int(staircaseSize); i++ {
			Hashprint := "%" + strconv.Itoa(int(staircaseSize)) + "s\n"
			fmt.Printf(Hashprint, strings.Repeat("#", i))
		}
	} else {
		log.Fatal("Invalid data for staircase size: ", staircaseSize)
	}

	return
}

func main() {

	createStaircase(8)

}
