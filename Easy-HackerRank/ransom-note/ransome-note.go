// Link to the problem statement: https://www.hackerrank.com/challenges/ctci-ransom-note/problem
// This is the "ransome note" solution implemented using an "array/slice" data structures.
// All test cases passed on hackerrank.com
package main

import (
	"fmt"
	"log"
)

var (
	validity                             bool
	canReplicateRansomeNoteFromMagazine  bool
	wordInRansomeNoteIsPresentInMagazine bool
)

func CheckDataValidity(magazine, note []string) bool {
	if len(magazine) >= 1 && len(magazine) <= 30000 && len(note) >= 1 && len(note) <= 30000 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data")
	}
	return validity
}

func CheckMagazine(magazine, note []string) {
	validity = CheckDataValidity(magazine, note)
	if validity == true {
		for i := range note {
			for j := range magazine {
				if note[i] == magazine[j] {
					wordInRansomeNoteIsPresentInMagazine = true
					magazine = append(magazine[:j], magazine[j+1:]...)
					break
				}
			}
			if wordInRansomeNoteIsPresentInMagazine == true {
				canReplicateRansomeNoteFromMagazine = true
			} else {
				canReplicateRansomeNoteFromMagazine = false
				break
			}
			wordInRansomeNoteIsPresentInMagazine = false
		}
		if canReplicateRansomeNoteFromMagazine == true {
			fmt.Println("Yes")
		} else if canReplicateRansomeNoteFromMagazine == false {
			fmt.Println("No")
		}
	}
}

func main() {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}

	CheckMagazine(magazine, note)
}
