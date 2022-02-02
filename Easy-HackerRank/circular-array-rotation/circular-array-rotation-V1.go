package main

import (
	"fmt"
	"log"
	"math"
)

var (
	k                     int32
	validity              bool
	queriesResultList     []int32
	valueOfElementAtIndex int32
)

func CheckDataValidity(a []int32, k int32, queries []int32) bool {
	arrayLengthOfA := len(a)
	arrayLengthOfQueries := len(queries)
	basePower5 := int32(math.Pow(10, 5))
	if arrayLengthOfA >= 1 && arrayLengthOfA <= int(basePower5) && arrayLengthOfQueries >= 1 &&
		arrayLengthOfQueries <= 500 && k >= 1 && k <= basePower5 {
		for i := range a {
			if a[i] >= 1 && a[i] <= basePower5 {
				validity = true
			} else {
				log.Fatalf("Invalid a array data: %d\n", a[i])
			}
		}

		for i := range queries {
			if queries[i] >= 0 && queries[i] < int32(arrayLengthOfA) {
				validity = true
			} else {
				log.Fatalf("Invalid queries array data: %d\n", queries[i])
			}
		}
	} else {
		log.Fatal("Invalid Data. Check the data.")
	}
	return validity
}

func main() {
	a := []int32{1, 2, 3}
	k = 2
	queries := []int32{0, 1, 2}
	arrayLengthOfQueries := len(queries)
	queriesResultList := make([]int32, arrayLengthOfQueries)
	validity = CheckDataValidity(a, k, queries)
	// Note:After the 2 rotations, element of index 0 which is (1) will be replaced by element of index 1 which is (2)
	// , element of index 1 which is (2) will be replaced by element of index 2 which is (3), and element of index
	// 2 which is (3) will be replaced by element of index 0 which is (1)
	if validity == true {
		for i := 0; i < len(queries); i++ {
			indexOfReplacingElementCalculation := queries[i] - k
			isNegativeNumber := math.Signbit(float64(indexOfReplacingElementCalculation))
			if isNegativeNumber == true {
				indexOfReplacingElementAfterKRotations := len(a) + int(indexOfReplacingElementCalculation)
				valueOfElementAtIndex = a[indexOfReplacingElementAfterKRotations]
				queriesResultList[i] = valueOfElementAtIndex
			} else {
				indexOfReplacingElementAfterKRotations := indexOfReplacingElementCalculation
				valueOfElementAtIndex = a[indexOfReplacingElementAfterKRotations]
				queriesResultList[i] = valueOfElementAtIndex
			}
		}
		fmt.Println(queriesResultList)
	}
}
