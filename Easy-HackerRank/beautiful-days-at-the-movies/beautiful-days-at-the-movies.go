// Link to problem statement: https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	i                     int32
	j                     int32
	k                     int32
	validity              bool
	numberOfBeautifulDays int32
)

func CheckDataValidity(i, j, k int32) bool {
	basePower6 := int32(math.Pow(10, 6))
	basePower6Times2 := basePower6 * 2
	basePower9 := int32(math.Pow(10, 9))
	basePower9Times2 := basePower9 * 2
	if i >= 1 && j >= i && j <= basePower6Times2 && k >= 1 && k <= basePower9Times2 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func findTheReverseOfANumber(number int32) int32 {
	stringReprOfTheNumber := strconv.Itoa(int(number))
	arrayForStringReprOfTheNumber := strings.Split(stringReprOfTheNumber, "")
	var newArray []string
	for i := len(arrayForStringReprOfTheNumber) - 1; i >= 0; i-- {
		newArray = append(newArray, arrayForStringReprOfTheNumber[i])
	}
	stringValueOfNewArray := strings.Join(newArray, "")
	intValueOfNewArray, err := strconv.Atoi(stringValueOfNewArray)
	if err != nil {
		log.Fatal(err)
	}
	return int32(intValueOfNewArray)
}

func BeautifulDays(i, j, k int32) int32 {
	validity = CheckDataValidity(i, j, k)
	numberOfBeautifulDays = 0
	if validity == true {
		startDay := i
		endDay := j
		for x := startDay; x <= endDay; x++ {
			reverseOfANumber := findTheReverseOfANumber(x)
			differenceBtwNumberAndItsReverse := int32(math.Abs(float64(x - reverseOfANumber)))
			if differenceBtwNumberAndItsReverse%k == 0 {
				numberOfBeautifulDays++
			}
		}
	}
	return numberOfBeautifulDays
}

func main() {
	i = 20
	j = 23
	k = 6
	numberOfBeautifulDays = BeautifulDays(i, j, k)
	fmt.Println(numberOfBeautifulDays)
}
