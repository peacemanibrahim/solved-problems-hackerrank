// Link to problem statement: https://www.hackerrank.com/challenges/bon-appetit/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	k        int32
	b        int32
	validity bool
)

func CheckDataValidity(bill []int32, k int32, b int32) bool {
	arrayLength := len(bill)
	basePower5 := math.Pow(10, 5)
	basePower4 := math.Pow(10, 4)
	billSumTotal := arraySum(bill)
	if arrayLength >= 2 && arrayLength <= int(basePower5) && k >= 0 && k < int32(arrayLength) {
		for i := range bill {
			if bill[i] >= 0 && bill[i] <= int32(basePower4) && b >= 0 && b <= billSumTotal {
				validity = true
			} else {
				log.Fatalf("Invalid array data: %d\n", bill[i])
			}
		}
	} else {
		log.Fatal("Invalid data for array length and/or k.")
	}
	return validity
}

func arraySum(arr []int32) int32 {
	var sum int32
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func BonAppetit(bill []int32, k int32, b int32) {
	validity = CheckDataValidity(bill, k, b)
	if validity == true {
		billSumTotal := arraySum(bill)
		costOfItemAnnaDidNotEat := bill[k]
		billToBeShared := billSumTotal - costOfItemAnnaDidNotEat
		costPerPerson := billToBeShared / 2
		amountAnnaOverpaid := math.Abs(float64(b - costPerPerson))
		intValueOfAmountAnnaOverpaid := int32(amountAnnaOverpaid)
		if intValueOfAmountAnnaOverpaid == 0 {
			fmt.Println("Bon Appetit")
		} else {
			fmt.Println(intValueOfAmountAnnaOverpaid)
		}
	}
}

func main() {
	bill := []int32{3, 10, 2, 9}
	k = 1
	b = 12

	BonAppetit(bill, k, b)
}

// Note: k is the index of the item that Anna declines to eat.
// Note: b is the amount that Anna paid
