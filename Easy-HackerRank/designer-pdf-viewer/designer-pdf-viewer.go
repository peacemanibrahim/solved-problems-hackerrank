// Link to problem statement: https://www.hackerrank.com/challenges/designer-pdf-viewer/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	validity               bool
	wordIndexes            []int32
	tallestCharacterHeight int32
	theHighlightedArea     int32
)

func CheckDataValidity(h []int32, word string) bool {
	wordArray := strings.Split(word, "")
	if len(wordArray) <= 10 {
		for i := range h {
			if h[i] >= 1 && h[i] <= 7 {
				validity = true
			} else {
				log.Fatalf("Invalid height array data: %d\n", h[i])
			}
		}
	} else {
		log.Fatal("Invalid word array length. Check the data.")
	}
	return validity
}

func findTallestCharacterHeight(h, wordIndexes []int32) int32 {
	tallestCharacterHeight = 0
	for i := range wordIndexes {
		index := wordIndexes[i]
		currentCharacterHeight := h[index]
		if currentCharacterHeight > tallestCharacterHeight {
			tallestCharacterHeight = currentCharacterHeight
		}
	}
	return tallestCharacterHeight
}

func DesignerPdfViewer(h []int32, word string) int32 {
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z"}
	wordArray := strings.Split(word, "")
	lengthOfWordArray := len(wordArray)
	wordIndexes := make([]int32, len(wordArray))
	validity = CheckDataValidity(h, word)
	if validity == true {
		for i := range wordArray {
			for j := range alphabets {
				if wordArray[i] == alphabets[j] {
					wordIndexes[i] = int32(j)
				}
			}
		}
		tallestCharacterHeight = findTallestCharacterHeight(h, wordIndexes)
		theHighlightedArea = tallestCharacterHeight * int32(lengthOfWordArray)
	}
	return theHighlightedArea
}

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	word := "abc"

	theHighlightedArea = DesignerPdfViewer(h, word)
	fmt.Println(theHighlightedArea)
}
