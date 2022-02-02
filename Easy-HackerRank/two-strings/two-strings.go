// Link to problem statement on hackerrank.com: https://www.hackerrank.com/challenges/two-strings/problem
// This is the solution implemented using "maps" data structure.
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	validity                      bool
	doStringsShareCommonSubstring string
)

func CheckDataValidity(s1, s2 string) bool {
	arrayRepOfStringS1 := strings.Split(s1, "")
	arrayRepOfStringS2 := strings.Split(s2, "")
	if len(arrayRepOfStringS1) >= 1 && len(arrayRepOfStringS1) <= 100000 && len(arrayRepOfStringS2) >= 1 &&
		len(arrayRepOfStringS2) <= 100000 {
		validity = true
	} else {
		log.Fatal("Invalid length of s1 and/or s2. Check the data.")
	}
	return validity
}

func TwoStrings(s1, s2 string) string {
	arrayRepOfStringS1 := strings.Split(s1, "")
	arrayRepOfStringS2 := strings.Split(s2, "")
	doStringsShareCommonSubstring = "No"

	validity = CheckDataValidity(s1, s2)
	if validity == true {
		mapOfS2 := make(map[string]string)
		for _, s := range arrayRepOfStringS2 {
			mapOfS2[s] = s
		}
		for i := range arrayRepOfStringS1 {
			keyOfSubstringToCheckForInS2 := arrayRepOfStringS1[i]
			substringInS2 := mapOfS2[keyOfSubstringToCheckForInS2]
			if substringInS2 != "" {
				doStringsShareCommonSubstring = "Yes"
				break
			}
		}
	}
	return doStringsShareCommonSubstring
}

func main() {
	s1 := "hello"
	s2 := "world"

	doStringsShareCommonSubstring = TwoStrings(s1, s2)
	fmt.Println(doStringsShareCommonSubstring)
}
