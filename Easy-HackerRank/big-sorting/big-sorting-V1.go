// THIS IS BUBBLE SORT: All test cases did Not pass on hackerrank
package main

import (
	"fmt"
	"log"
	"math/big"
)

var (
	validity      bool
	arrayIsSorted bool
)

func CheckDataValidity(unsorted []string) bool {
	if len(unsorted) >= 1 && len(unsorted) <= 200000 {
		validity = true
	} else {
		log.Fatal("Invalid unsorted array length. Check the data.")
	}
	return validity
}

func BigSorting(unsorted []string) []string {
	arrayIsSorted = false
	validity = CheckDataValidity(unsorted)
	if validity == true {
		for arrayIsSorted == false {
			arrayIsSorted = true
			for i := 0; i < len(unsorted)-1; i++ {
				bigIntRepOfNumberString, _ := new(big.Int).SetString(unsorted[i], 0)
				bigIntRepOfNumberAfterNumberString, _ := new(big.Int).SetString(unsorted[i+1], 0)
				resultOfCheckingWhichNumberIsGreater := bigIntRepOfNumberString.Cmp(bigIntRepOfNumberAfterNumberString)
				if resultOfCheckingWhichNumberIsGreater == 1 {
					theStoredValueOfIndexI := unsorted[i]
					unsorted[i] = unsorted[i+1]
					unsorted[i+1] = theStoredValueOfIndexI
					arrayIsSorted = false
				}
			}
		}
	}
	return unsorted
}

func main() {
	unsorted := []string{"100", "6046724675206781253805417697364063874537349942332939012160502967377173420319837265456476",
		"72091970412420810", "3746855980099062683149629582771", "4763506513", "4381858583387279277494367930538",
		"195294541694873319042720380485249001380825565396311663127260228626400961443707964165554187168383846", "57500297590012603652986133599394871645776460",
		"3443449340302668605052308", "799898013447209990576845913871859165418584121624031762316631"}

	sortedArray := BigSorting(unsorted)
	fmt.Println(sortedArray)
}

// Note: Use the  func (x *Int) Cmp(y *Int) (r int) in math/big package to compare which one is bigger in two
// big integer numbers. This is because you cannot use logical operators (<, >, =, etc.) on big integers.
