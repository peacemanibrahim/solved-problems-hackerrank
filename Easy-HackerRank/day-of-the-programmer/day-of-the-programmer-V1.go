package main

import (
	"fmt"
	"log"
)

var (
	year                        int32
	validity                    bool
	isLeapYear                  bool
	officialCalendarDuringAYear string
)

func CheckDataValidity(year int32) bool {
	if year >= 1700 && year <= 2700 {
		validity = true
	} else {
		log.Fatal("Invalid year entered. Check the year.")
	}
	return validity
}

func CheckOfficialRussianCalendarDuringAYear(year int32) string {
	if year >= 1700 && year <= 1917 {
		officialCalendarDuringAYear = "Julian calendar"
	} else if year == 1918 {
		officialCalendarDuringAYear = "Transitioning year"
	} else {
		officialCalendarDuringAYear = "Gregorian calendar"
	}
	return officialCalendarDuringAYear
}

func main() {
	year = 1918
	validity = CheckDataValidity(year)
	officialCalendarDuringAYear = CheckOfficialRussianCalendarDuringAYear(year)

	if validity == true {
		if officialCalendarDuringAYear == "Julian calendar" {
			yearDividedBy4Remainder := year % 4
			if yearDividedBy4Remainder == 0 {
				isLeapYear = true
				fmt.Printf("12-09-%d\n", year)
			} else {
				isLeapYear = false
				fmt.Printf("13-09-%d\n", year)
			}
		} else if officialCalendarDuringAYear == "Gregorian calendar" {
			yearDividedBy400Remainder := year % 400
			yearDividedBy4Remainder := year % 4
			yearDividedBy100Remainder := year % 100
			if yearDividedBy400Remainder == 0 {
				isLeapYear = true
				fmt.Printf("12-09-%d\n", year)
			} else if yearDividedBy4Remainder == 0 && yearDividedBy100Remainder != 0 {
				isLeapYear = true
				fmt.Printf("12-09-%d\n", year)
			} else {
				isLeapYear = false
				fmt.Printf("13-09-%d\n", year)
			}
		} else {
			isLeapYear = false
			fmt.Printf("26-09-%d\n", year)
		}
	}
}
