package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	p        int32
	q        int32
	validity bool
)

func CheckDataValidity(p, q int32) bool {
	if p > 0 && q > p && q < 100000 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func GetLeftPartOfSquareOfNumber(arrayRepOfSquareOfNumber []string, leftPartLength int) int {
	var arrayOfLeftPart []string
	for i := 0; i < leftPartLength; i++ {
		arrayOfLeftPart = append(arrayOfLeftPart, arrayRepOfSquareOfNumber[i])
	}
	arrayOfLeftPartConvertedToString := strings.Join(arrayOfLeftPart, "")
	intRepOfString, _ := strconv.Atoi(arrayOfLeftPartConvertedToString)
	return intRepOfString
}

func GetRightPartOfSquareOfNumber(arrayRepOfSquareOfNumber []string, leftPartLength, lengthOfArrayRepOfSquareOfNumber int) int {
	var arrayOfRightPart []string
	for i := leftPartLength; i < lengthOfArrayRepOfSquareOfNumber; i++ {
		arrayOfRightPart = append(arrayOfRightPart, arrayRepOfSquareOfNumber[i])
	}
	arrayOfRightPartConvertedToString := strings.Join(arrayOfRightPart, "")
	intRepOfString, _ := strconv.Atoi(arrayOfRightPartConvertedToString)
	return intRepOfString
}

func main() {
	p = 1
	q = 100
	validity = CheckDataValidity(p, q)

	if validity == true {
		for i := p; i <= q; i++ {
			strRepOfNumber := strconv.Itoa(int(i))
			arrayRepOfNumber := strings.Split(strRepOfNumber, "")
			d := len(arrayRepOfNumber)

			squareOfNumber := i * i
			strRepOfSquareOfNumber := strconv.Itoa(int(squareOfNumber))
			arrayRepOfSquareOfNumber := strings.Split(strRepOfSquareOfNumber, "")
			lengthOfArrayRepOfSquareOfNumber := len(arrayRepOfSquareOfNumber)

			lengthOfLeftPartOfSquareOfNumber := lengthOfArrayRepOfSquareOfNumber - d
			// lengthOfRightPartOfSquareOfNumber := d

			leftPartAfterSplittingTheSquare := GetLeftPartOfSquareOfNumber(arrayRepOfSquareOfNumber, lengthOfLeftPartOfSquareOfNumber)
			rightPartAfterSplittingTheSquare := GetRightPartOfSquareOfNumber(arrayRepOfSquareOfNumber, lengthOfLeftPartOfSquareOfNumber, lengthOfArrayRepOfSquareOfNumber)
			sumOfLeftAndRightPartOfTheSquare := leftPartAfterSplittingTheSquare + rightPartAfterSplittingTheSquare
			if sumOfLeftAndRightPartOfTheSquare == int(i) {
				fmt.Println(i)
			}
		}
	}
}

// Note: d is the original length (before being squared) of each number in the incluse range of p and q.
