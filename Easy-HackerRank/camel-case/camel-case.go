// Link to problem statement: https://www.hackerrank.com/challenges/camelcase/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	validity            bool
	capitalLettersCount int32
	noOfWordsInS        int32
)

func CheckDataValidity(s string) bool {
	arrayRepOfS := strings.Split(s, "")
	if len(arrayRepOfS) >= 1 && len(arrayRepOfS) <= 100000 {
		validity = true
	} else {
		log.Fatal("Invalid length of s. Check the data.")
	}
	return validity
}

func CamelCase(s string) int32 {
	arrayRepOfS := strings.Split(s, "")
	alphabetsInCapital := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	capitalLettersCount = 0
	noOfWordsInS = 0

	validity = CheckDataValidity(s)
	if validity == true {
		if len(arrayRepOfS) == 1 {
			return 0
		}
		for i := 0; i < len(arrayRepOfS); i++ {
			for j := 0; j < len(alphabetsInCapital); j++ {
				if arrayRepOfS[i] == alphabetsInCapital[j] {
					capitalLettersCount++
					break
				}
			}
		}
		if capitalLettersCount == 0 && len(arrayRepOfS) > 1 {
			noOfWordsInS = 1
		} else {
			noOfWordsInS = capitalLettersCount + 1
		}
	}
	return noOfWordsInS
}

func main() {
	s := "oneTwoThree"

	noOfWordsInS = CamelCase(s)
	fmt.Println(noOfWordsInS)
}
