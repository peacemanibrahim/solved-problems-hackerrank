// Link to problem statement: https://www.hackerrank.com/challenges/library-fine/problem
package main

import (
	"fmt"
	"log"
)

var (
	d1        int32
	m1        int32
	y1        int32
	d2        int32
	m2        int32
	y2        int32
	validity  bool
	fineToPay int32
)

func CheckDataValidity(d1, m1, y1, d2, m2, y2 int32) bool {
	if d1 >= 1 && d1 <= 31 && d2 >= 1 && d2 <= 31 && m1 >= 1 && m1 <= 12 && m2 >= 1 && m2 <= 12 && y1 >= 1 &&
		y1 <= 3000 && y2 >= 1 && y2 <= 3000 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func LibraryFine(d1, m1, y1, d2, m2, y2 int32) int32 {
	validity = CheckDataValidity(d1, m1, y1, d2, m2, y2)
	if validity == true {
		if y1 == y2 && m1 == m2 && d1 <= d2 {
			fineToPay = 0
		} else if y1 == y2 && m1 == m2 && d1 > d2 {
			numberOfDaysLate := d1 - d2
			fineToPay = 15 * numberOfDaysLate
		} else if y1 == y2 && m1 > m2 {
			numberOfMonthsLate := m1 - m2
			fineToPay = 500 * numberOfMonthsLate
		} else if y1 > y2 {
			fineToPay = 10000
		}
	}
	return fineToPay
}

func main() {
	d1 = 9
	m1 = 6
	y1 = 2015
	d2 = 6
	m2 = 6
	y2 = 2015

	fineToPay = LibraryFine(d1, m1, y1, d2, m2, y2)
	fmt.Println(fineToPay)
}

// Note: d1, m1, y1 is the return date for the library book
// Note: d2, m2, y2 is due date of return for the library book
