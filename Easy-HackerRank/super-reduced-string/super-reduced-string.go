// Link to problem statement: https://www.hackerrank.com/challenges/reduced-string/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	validity                    bool
	resultingSuperReducedString string
)

func CheckDataValidity(s string) bool {
	arrayRepOfS := strings.Split(s, "")
	if len(arrayRepOfS) >= 1 && len(arrayRepOfS) <= 100 {
		validity = true
	} else {
		log.Fatal("Invalid length of s. Check the data.")
	}
	return validity
}

func SuperReducedString(s string) string {
	arrayRepOfS := strings.Split(s, "")
	validity = CheckDataValidity(s)
	if validity == true {
		i := 1
		for i < len(arrayRepOfS)-1 && len(arrayRepOfS) > 2 {
			if i > 0 && arrayRepOfS[i] == arrayRepOfS[i-1] {
				arrayRepOfS = append(arrayRepOfS[:i-1], arrayRepOfS[i+1:]...)
				i--
			} else if arrayRepOfS[i] == arrayRepOfS[i+1] {
				arrayRepOfS = append(arrayRepOfS[:i], arrayRepOfS[i+2:]...)
			} else {
				i++
			}
		}

		if len(arrayRepOfS) == 2 {
			if arrayRepOfS[0] == arrayRepOfS[1] {
				return "Empty String"
			} else {
				resultingSuperReducedString := strings.Join(arrayRepOfS, "")
				return resultingSuperReducedString
			}
		} else if len(arrayRepOfS) == 1 {
			resultingSuperReducedString = strings.Join(arrayRepOfS, "")
			return resultingSuperReducedString
		} else {
			resultingSuperReducedString = strings.Join(arrayRepOfS, "")
		}
	}
	return resultingSuperReducedString
}

func main() {
	s := "aaabccddd"

	resultingSuperReducedString = SuperReducedString(s)
	fmt.Println(resultingSuperReducedString)
}
