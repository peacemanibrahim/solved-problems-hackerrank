// Link to problem statement: https://www.hackerrank.com/challenges/angry-professor/problem
package main

import (
	"fmt"
	"log"
)

var (
	k                            int32
	validity                     bool
	numberOfStudentInClassONTime int32
	isTheClassCanceled           string
)

func CheckDataValidity(k int32, a []int32) bool {
	arrayLength := len(a)
	if arrayLength >= 1 && arrayLength <= 1000 && k >= 1 && k <= int32(arrayLength) {
		for i := range a {
			if a[i] >= -100 && a[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", a[i])
			}
		}
	} else {
		log.Fatal("Invalid array length and/or k. Check the data.")
	}
	return validity
}

func AngryProfessor(k int32, a []int32) string {
	numberOfStudentInClassONTime = 0
	validity = CheckDataValidity(k, a)
	if validity == true {
		for i := range a {
			if a[i] <= 0 {
				numberOfStudentInClassONTime++
			}
			if numberOfStudentInClassONTime >= k {
				isTheClassCanceled = "NO"
				break
			} else {
				isTheClassCanceled = "YES"
			}
		}
	}
	return isTheClassCanceled
}

func main() {
	k = 2
	a := []int32{0, -1, 2, 1}

	isTheClassCanceled = AngryProfessor(k, a)
	fmt.Println(isTheClassCanceled)
}
