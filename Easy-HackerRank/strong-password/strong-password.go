// Link to problem statement: https://www.hackerrank.com/challenges/strong-password/problem
package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	n                            int32
	validity                     bool
	necessaryCharacterTypesCount int32
	minimumNoOfCharactersToAdd   int32
	shouldBreak                  bool
)

const (
	requiredPasswordLength              = 6
	requiredNoOfNecessaryCharacterTypes = 4
)

func CheckDataValidity(password string) bool {
	arrayRepOfPassword := strings.Split(password, "")
	if len(arrayRepOfPassword) >= 1 && len(arrayRepOfPassword) <= 100 {
		validity = true
	} else {
		log.Fatal("Invalid password length. Check the data.")
	}
	return validity
}

func CheckIfNumbersArePresentInPassword(password string) {
	shouldBreak = false
	arrayRepOfPassword := strings.Split(password, "")
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := range arrayRepOfPassword {
		for j := range numbers {
			if arrayRepOfPassword[i] == numbers[j] {
				necessaryCharacterTypesCount++
				shouldBreak = true
				break
			}
		}
		if shouldBreak == true {
			break
		}
	}
}

func CheckIfLowerCaseLettersArePresentInPassword(password string) {
	shouldBreak = false
	arrayRepOfPassword := strings.Split(password, "")
	lowerCaseLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := range arrayRepOfPassword {
		for j := range lowerCaseLetters {
			if arrayRepOfPassword[i] == lowerCaseLetters[j] {
				necessaryCharacterTypesCount++
				shouldBreak = true
				break
			}
		}
		if shouldBreak == true {
			break
		}
	}
}

func CheckIfUpperCaseLettersArePresentInPassword(password string) {
	shouldBreak = false
	arrayRepOfPassword := strings.Split(password, "")
	upperCaseLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := range arrayRepOfPassword {
		for j := range upperCaseLetters {
			if arrayRepOfPassword[i] == upperCaseLetters[j] {
				necessaryCharacterTypesCount++
				shouldBreak = true
				break
			}
		}
		if shouldBreak == true {
			break
		}
	}
}

func CheckIfSpecialCharactersArePresentInPassword(password string) {
	shouldBreak = false
	arrayRepOfPassword := strings.Split(password, "")
	specialCharacter := []string{"!", "@", "#", "$", "%", "^", "&", "*", "()", "-", "+"}
	for i := range arrayRepOfPassword {
		for j := range specialCharacter {
			if arrayRepOfPassword[i] == specialCharacter[j] {
				necessaryCharacterTypesCount++
				shouldBreak = true
				break
			}
		}
		if shouldBreak == true {
			break
		}
	}
}

func MinimumNumber(n int32, password string) int32 {
	arrayRepOfPassword := strings.Split(password, "")
	validity = CheckDataValidity(password)
	if validity == true {
		CheckIfNumbersArePresentInPassword(password)
		CheckIfLowerCaseLettersArePresentInPassword(password)
		CheckIfUpperCaseLettersArePresentInPassword(password)
		CheckIfSpecialCharactersArePresentInPassword(password)

		noOfNecessaryCharacterTypesNotPresentInPassword := requiredNoOfNecessaryCharacterTypes -
			necessaryCharacterTypesCount
		if len(arrayRepOfPassword) >= requiredPasswordLength {
			minimumNoOfCharactersToAdd = noOfNecessaryCharacterTypesNotPresentInPassword
		} else if len(arrayRepOfPassword) < requiredPasswordLength {
			additionOfPasswordLengthAndNoOfNecessaryCharactersNotPresent := len(arrayRepOfPassword) +
				int(noOfNecessaryCharacterTypesNotPresentInPassword)
			if additionOfPasswordLengthAndNoOfNecessaryCharactersNotPresent >= requiredPasswordLength {
				minimumNoOfCharactersToAdd = noOfNecessaryCharacterTypesNotPresentInPassword
			} else {
				noOfCharactersToMakePasswordMeetRequiredLength := requiredPasswordLength -
					additionOfPasswordLengthAndNoOfNecessaryCharactersNotPresent
				minimumNoOfCharactersToAdd = noOfNecessaryCharacterTypesNotPresentInPassword +
					int32(noOfCharactersToMakePasswordMeetRequiredLength)
			}
		}
	}
	return minimumNoOfCharactersToAdd
}

func main() {
	n = 11
	password := "#HackerRank"

	minimumNoOfCharactersToAdd = MinimumNumber(n, password)
	fmt.Println(minimumNoOfCharactersToAdd)
}
