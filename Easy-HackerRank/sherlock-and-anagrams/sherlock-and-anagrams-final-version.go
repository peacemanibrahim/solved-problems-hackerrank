// Link to problem statement: https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
// This is the sherlock and anagrams solution implemented using "maps" data structure
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	validity                              bool
	substringsAreAnagram                  bool
	pairsOfSubstringsThatAreAnagramsCount int32
)

func CheckDataValidity(s string) bool {
	arrayRepOfS := strings.Split(s, "")
	if len(arrayRepOfS) >= 2 && len(arrayRepOfS) <= 100 {
		validity = true
	} else {
		log.Fatal("Invalid length of s. Check the data.")
	}
	return validity
}

func checkIfAPairOfSubstringsAreAnagrams(s1, s2 []string) bool {
	substringsAreAnagram = false
	mapOfS1 := make(map[string]int)
	for _, s := range s1 {
		mapOfS1[s]++
	}
	mapOfS2 := make(map[string]int)
	for _, s := range s2 {
		mapOfS2[s]++
	}

	for i := range s1 {
		keyOfSubstringToCheckFor := s1[i]
		noOfOccurrenceOfSubstringInS2, ok := mapOfS2[keyOfSubstringToCheckFor]
		noOfOccurrenceOfSubstringInS1 := mapOfS1[keyOfSubstringToCheckFor]
		if ok == true && noOfOccurrenceOfSubstringInS1 == noOfOccurrenceOfSubstringInS2 {
			substringsAreAnagram = true
		} else {
			substringsAreAnagram = false
			break
		}
	}
	return substringsAreAnagram
}

func SherlockAndAnagrams(s string) int32 {
	arrayRepOfS := strings.Split(s, "")
	newArrayRepOfS := arrayRepOfS
	s1 := []string{}
	s2 := []string{}

	validity = CheckDataValidity(s)
	if validity == true {
		for b := 1; b < len(arrayRepOfS); b++ {
			i := b
			for (i + 1) <= len(newArrayRepOfS) {
				s1 = newArrayRepOfS[:i]
				for j := 1; (j + i) <= len(newArrayRepOfS); j++ {
					s2 = newArrayRepOfS[j : j+i]
					substringsAreAnagram = checkIfAPairOfSubstringsAreAnagrams(s1, s2)
					if substringsAreAnagram == true {
						pairsOfSubstringsThatAreAnagramsCount++
						substringsAreAnagram = false
					}
					s2 = []string{}
				}
				s1 = []string{}
				newArrayRepOfS = newArrayRepOfS[1:]
			}
			newArrayRepOfS = arrayRepOfS
		}
	}
	return pairsOfSubstringsThatAreAnagramsCount
}

func main() {
	s := "bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc"

	result := SherlockAndAnagrams(s)
	fmt.Println(result)
}
