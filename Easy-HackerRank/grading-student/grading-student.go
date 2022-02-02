// Link to problem statement: https://www.hackerrank.com/challenges/grading/problem
package main

import (
	"fmt"
	"log"
)

var (
	validity bool
)

const (
	passingGrade int32 = 38
	multiple     int32 = 5
)

func CheckDataValidity(grades []int32) bool {
	arrayLength := len(grades)
	if arrayLength >= 1 && arrayLength <= 60 {
		for i := range grades {
			if grades[i] >= 0 && grades[i] <= 100 {
				validity = true
			} else {
				log.Fatalf("Invalid Array Data %d\n", grades[i])
			}
		}
	} else {
		log.Fatal("Invalid Array Length")
	}
	return validity
}

func FindNearestMultiplesOfANumber(number int32) int32 {
	multiplyingFactor := number / multiple
	nearestMultiple := (multiplyingFactor * multiple) + multiple
	return nearestMultiple
}

func GradingStudents(grades []int32) []int32 {
	var FinalGrade []int32
	arrayLength := len(grades)
	validity = CheckDataValidity(grades)
	if validity == true {
		for i := 0; i < arrayLength; i++ {
			if grades[i] >= passingGrade {
				nearestMultipleOfGrade := FindNearestMultiplesOfANumber(grades[i])
				nearestMultipleAndGradeDifference := nearestMultipleOfGrade - grades[i]
				if nearestMultipleAndGradeDifference < 3 {
					FinalGrade = append(FinalGrade, nearestMultipleOfGrade)
				} else {
					FinalGrade = append(FinalGrade, grades[i])
				}
			} else {
				FinalGrade = append(FinalGrade, grades[i])
			}
		}
	}
	return FinalGrade
}

func main() {
	grades := []int32{73, 67, 38, 33}
	FinalGrade := GradingStudents(grades)
	fmt.Println(FinalGrade)
}
