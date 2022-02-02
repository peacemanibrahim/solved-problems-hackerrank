// Link to problem statement: https://www.hackerrank.com/challenges/two-characters/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	l                                                      int32
	validity                                               bool
	shouldBreak                                            bool
	listOfDistinctLettersInS                               []string
	precedingAlternatingLetterInString                     string
	alternatingLettersInStringCount                        int32
	lengthOfLongestStringPossibleWithTwoAlternatingLetters int32
)

func CheckDataValidity(s string) bool {
	arrayRepOfS := strings.Split(s, "")
	if len(arrayRepOfS) >= 1 && len(arrayRepOfS) <= 1000 {
		validity = true
	} else {
		log.Fatal("Invalid length of s. Check the data.")
	}
	return validity
}

func GetListOfDistinctLettersInS(s string) []string {
	var letterAlreadyOccurredOnceInS bool = false
	arrayRepOfS := strings.Split(s, "")
	listOfDistinctLettersInS = []string{}
	listOfDistinctLettersInS = append(listOfDistinctLettersInS, arrayRepOfS[0])
	for i := range arrayRepOfS {
		for j := range listOfDistinctLettersInS {
			if arrayRepOfS[i] == listOfDistinctLettersInS[j] {
				letterAlreadyOccurredOnceInS = true
				break
			}
		}
		if letterAlreadyOccurredOnceInS == false {
			listOfDistinctLettersInS = append(listOfDistinctLettersInS, arrayRepOfS[i])
		}
		letterAlreadyOccurredOnceInS = false
	}
	return listOfDistinctLettersInS
}

func Alternate(l int32, s string) int32 {
	arrayRepOfS := strings.Split(s, "")
	alternatingLettersInStringCount = 1
	lengthOfLongestStringPossibleWithTwoAlternatingLetters = 0
	precedingAlternatingLetterInString = ""
	shouldBreak = false

	validity = CheckDataValidity(s)
	if validity == true {
		listOfDistinctLettersInS = GetListOfDistinctLettersInS(s)
		for len(listOfDistinctLettersInS) >= 2 {
			for i := 1; i < len(listOfDistinctLettersInS); i++ {
				for j := 0; j < len(arrayRepOfS); j++ {
					if listOfDistinctLettersInS[0] == arrayRepOfS[j] || listOfDistinctLettersInS[i] == arrayRepOfS[j] {
						if precedingAlternatingLetterInString != "" {
							if precedingAlternatingLetterInString == arrayRepOfS[j] {
								shouldBreak = true
								break
							} else {
								alternatingLettersInStringCount++
							}
						}
						precedingAlternatingLetterInString = arrayRepOfS[j]
					}
				}
				if shouldBreak == false {
					if alternatingLettersInStringCount > lengthOfLongestStringPossibleWithTwoAlternatingLetters {
						lengthOfLongestStringPossibleWithTwoAlternatingLetters = alternatingLettersInStringCount
					}
				}
				alternatingLettersInStringCount = 1
				precedingAlternatingLetterInString = ""
				shouldBreak = false
			}
			listOfDistinctLettersInS = listOfDistinctLettersInS[1:]
		}
	}
	return lengthOfLongestStringPossibleWithTwoAlternatingLetters
}

func main() {
	s := "beabeefeab"
	l = 10

	result := Alternate(l, s)
	fmt.Println(result)
}

// Note: l is the length of string s
