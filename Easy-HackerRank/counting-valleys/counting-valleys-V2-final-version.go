// Link to problem statement: https://www.hackerrank.com/challenges/counting-valleys/problem
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

func CheckDataValidity(pathArray []string) bool {
	basePower6 := int(math.Pow(10, 6))
	arrayLength := len(pathArray)
	if arrayLength >= 2 && arrayLength <= basePower6 {
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

func ConvertPathArrayToNumberRepresentation(pathArray []string) []int32 {
	arrayLength := len(pathArray)
	numberRepresentationOfPathArray := make([]int32, arrayLength)
	for i := range pathArray {
		if pathArray[i] == "D" {
			numberRepresentationOfPathArray[i] = -1
		} else {
			numberRepresentationOfPathArray[i] = 1
		}
	}
	return numberRepresentationOfPathArray
}

func CountingValleys(steps int32, path string) int32 {
	pathArray := strings.Split(path, "")
	validity = CheckDataValidity(pathArray)
	if validity == true {
		numberRepresentationOfPathArray = ConvertPathArrayToNumberRepresentation(pathArray)
		hikerPosition := 0
		valleyCount = 0
		seaLevel := 0
		for i := range numberRepresentationOfPathArray {
			hikerPosition += int(numberRepresentationOfPathArray[i])
			if hikerPosition <= -1 {
				isTrueValley = true
			} else if hikerPosition >= 1 {
				isTrueValley = false
			}
			if isTrueValley == true && hikerPosition == seaLevel {
				valleyCount++
			}
		}
	}
	return valleyCount
}

func main() {
	path = "UUUUDDDUDDDDUUDDUU"
	steps = 14

	valleyCount = CountingValleys(steps, path)
	fmt.Println(valleyCount)
}
