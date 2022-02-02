package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

var (
	steps                           int32
	path                            string
	validity                        bool
	numberRepresentationOfPathArray []int32
	valleyCount                     int32
	isTrueValley                    bool
)

func CheckDataValidity(steps int32, pathArray []string) bool {
	basePower6 := math.Pow(10, 6)
	if steps >= 2 && steps <= int32(basePower6) {
		for i := range pathArray {
			if pathArray[i] == "D" || pathArray[i] == "U" {
				validity = true
			} else {
				log.Fatalf("Invalid array data %s\n", pathArray[i])
			}
		}
	} else {
		log.Fatal("Invalid data for steps. Check the data")
	}
	return validity
}

func main() {
	path = "UUUUDDDUDDDDUUDDUU"
	pathArray := strings.Split(path, "")
	steps = int32(len(pathArray))
	validity = CheckDataValidity(steps, pathArray)
	numberRepresentationOfPathArray := make([]int32, steps)

	if validity == true {
		for i := range pathArray {
			if pathArray[i] == "D" {
				numberRepresentationOfPathArray[i] = -1
			} else {
				numberRepresentationOfPathArray[i] = 1
			}
		}
		hikerPosition := 0
		valleyCount = 0
		seaLevel := 0
		for i := range numberRepresentationOfPathArray {
			hikerPosition += int(numberRepresentationOfPathArray[i])
			// The below is because I want to get the amount of times that the hiker came back to sea level
			// after going below sea level, this is in order to find the number of valleys the hiker walked through.
			// I do NOT want to get the amount of times that the hiker went above the sea level and came back
			// to the sealevel.
			if hikerPosition <= -1 {
				isTrueValley = true
			} else if hikerPosition >= 1 {
				isTrueValley = false
			}
			if isTrueValley == true && hikerPosition == seaLevel {
				valleyCount++
			}
		}
		fmt.Println(valleyCount)
	}
}
