// Link to problem statement: https://www.hackerrank.com/challenges/big-sorting/problem
// Merge Sort: This is merge sort using recursion. More test cases passed on hackerrank but Not all
package main

import (
	"fmt"
	"log"
	"math/big"
)

var (
	validity bool
)

func CheckDataValidity(unsorted []string) bool {
	if len(unsorted) >= 1 && len(unsorted) <= 200000 {
		validity = true
	} else {
		log.Fatal("Invalid unsorted array length. Check the data.")
	}
	return validity
}

func MergeLeftAndRightPartOfArrray(leftPartOfArray, rightPartOfArray []string) []string {
	sortedArrayOfNumberString := []string{}
	i := 0
	j := 0
	for i < len(leftPartOfArray) && j < len(rightPartOfArray) {
		bigIntRepOfNumberStringInLeftPartOfArray, _ := new(big.Int).SetString(leftPartOfArray[i], 0)
		bigIntRepOfNumberStringInRightPartOfArray, _ := new(big.Int).SetString(rightPartOfArray[j], 0)
		resultOfCheckingWhichNumberIsLesser :=
			bigIntRepOfNumberStringInLeftPartOfArray.Cmp(bigIntRepOfNumberStringInRightPartOfArray)
		if resultOfCheckingWhichNumberIsLesser == -1 {
			sortedArrayOfNumberString = append(sortedArrayOfNumberString, leftPartOfArray[i])
			i++
		} else {
			sortedArrayOfNumberString = append(sortedArrayOfNumberString, rightPartOfArray[j])
			j++
		}
	}
	for m := i; m < len(leftPartOfArray); m++ {
		sortedArrayOfNumberString = append(sortedArrayOfNumberString, leftPartOfArray[m])
	}
	for p := j; p < len(rightPartOfArray); p++ {
		sortedArrayOfNumberString = append(sortedArrayOfNumberString, rightPartOfArray[p])
	}
	return sortedArrayOfNumberString
}

func BigSorting(unsorted []string) []string {
	var leftPartOfArray []string
	var rightPartOfArray []string
	validity = CheckDataValidity(unsorted)
	if validity == true {
		arrayLength := len(unsorted)
		if arrayLength == 1 {
			return unsorted
		}
		midPointOfArray := arrayLength / 2
		leftPartOfArray = BigSorting(unsorted[:midPointOfArray])
		rightPartOfArray = BigSorting(unsorted[midPointOfArray:])
	}
	sortedArrayOfNumberString := MergeLeftAndRightPartOfArrray(leftPartOfArray, rightPartOfArray)
	return sortedArrayOfNumberString
}

func main() {
	unsorted := []string{"100", "6046724675206781253805417697364063874537349942332939012160502967377173420319837265456476",
		"72091970412420810", "3746855980099062683149629582771", "4763506513", "4381858583387279277494367930538",
		"195294541694873319042720380485249001380825565396311663127260228626400961443707964165554187168383846", "57500297590012603652986133599394871645776460",
		"3443449340302668605052308", "799898013447209990576845913871859165418584121624031762316631"}

	sortedArrayOfNumberString := BigSorting(unsorted)
	fmt.Println(sortedArrayOfNumberString)
}
