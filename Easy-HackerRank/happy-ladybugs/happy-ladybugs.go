// Link to problem statement: https://www.hackerrank.com/challenges/happy-ladybugs/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	n                int32
	validity         bool
	areLadyBugsHappy string
	underscoreCount  int32
	eachLetterCount  int32
)

func CheckDataValidity(b string) bool {
	arrayRepOfB := strings.Split(b, "")
	if len(arrayRepOfB) >= 1 && len(arrayRepOfB) <= 100 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func findIfLadybugsAreHappyWhenBHasOneCell(arrayRepOfB []string) string {
	if arrayRepOfB[0] == "_" {
		areLadyBugsHappy = "YES"
	} else {
		areLadyBugsHappy = "NO"
	}
	return areLadyBugsHappy
}

func findIfLadybugsAreHappyWhenBHasTwoCells(arrayRepOfB []string) string {
	if arrayRepOfB[0] == arrayRepOfB[1] {
		areLadyBugsHappy = "YES"
	} else {
		areLadyBugsHappy = "NO"
	}
	return areLadyBugsHappy
}

func checkIfStringBAlreadyHasHappyLadybugsByDefault(arrayRepOfB []string) string {
	var areLadyBugsHappyByDefault string
	for i := 1; i < len(arrayRepOfB)-1; i++ {
		if arrayRepOfB[len(arrayRepOfB)-1] == arrayRepOfB[len(arrayRepOfB)-2] {
			if arrayRepOfB[i-1] == arrayRepOfB[i] || arrayRepOfB[i+1] == arrayRepOfB[i] {
				areLadyBugsHappyByDefault = "YES"
			} else {
				areLadyBugsHappyByDefault = "NO"
				break
			}
		} else {
			areLadyBugsHappyByDefault = "NO"
		}
	}
	return areLadyBugsHappyByDefault
}

func findIfLadybugsAreHappyWhenBHasMoreThanTwoCells(arrayRepOfB, alphabets []string) string {
	underscoreCount = 0
	eachLetterCount = 0
	areLadyBugsHappyByDefault := checkIfStringBAlreadyHasHappyLadybugsByDefault(arrayRepOfB)
	areLadyBugsHappy = areLadyBugsHappyByDefault
	for i := 0; i < len(arrayRepOfB); i++ {
		if arrayRepOfB[i] == "_" {
			underscoreCount++
		}
	}
	if underscoreCount >= 1 {
		for i := range alphabets {
			for j := range arrayRepOfB {
				if alphabets[i] == arrayRepOfB[j] {
					eachLetterCount++
				}
				if eachLetterCount >= 2 {
					break
				}
			}
			if eachLetterCount == 1 {
				areLadyBugsHappy = "NO"
				break
			}
			areLadyBugsHappy = "YES"
			eachLetterCount = 0
		}
	}
	return areLadyBugsHappy
}

func HappyLadybugs(n int32, b string) string {
	arrayRepOfB := strings.Split(b, "")
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z"}
	validity = CheckDataValidity(b)
	if validity == true {
		if len(arrayRepOfB) == 1 {
			areLadyBugsHappy = findIfLadybugsAreHappyWhenBHasOneCell(arrayRepOfB)
		} else if len(arrayRepOfB) == 2 {
			areLadyBugsHappy = findIfLadybugsAreHappyWhenBHasTwoCells(arrayRepOfB)
		} else if len(arrayRepOfB) > 2 {
			areLadyBugsHappy = findIfLadybugsAreHappyWhenBHasMoreThanTwoCells(arrayRepOfB, alphabets)
		}
	}
	return areLadyBugsHappy
}

func main() {
	n = 9
	b := "AABBCVCV_"

	areLadyBugsHappy = HappyLadybugs(n, b)
	fmt.Println(areLadyBugsHappy)
}

// Note: arrayRepOfB means array representation of b string
