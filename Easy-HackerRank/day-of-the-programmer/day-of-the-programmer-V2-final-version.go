// Link to problem statement: https://www.hackerrank.com/challenges/day-of-the-programmer/problem
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
	dayOfProgrammer             string
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

func CheckJulianCalendarLeapYear(year int32) bool {
	yearDividedBy4Remainder := year % 4
	if yearDividedBy4Remainder == 0 {
		isLeapYear = true
	} else {
		isLeapYear = false
	}
	return isLeapYear
}

func CheckGregorianCalendarLeapYear(year int32) bool {
	yearDividedBy400Remainder := year % 400
	yearDividedBy4Remainder := year % 4
	yearDividedBy100Remainder := year % 100
	if yearDividedBy400Remainder == 0 {
		isLeapYear = true
	} else if yearDividedBy4Remainder == 0 && yearDividedBy100Remainder != 0 {
		isLeapYear = true
	} else {
		isLeapYear = false
	}
	return isLeapYear
}

func FindDayOfProgrammer(year int32) string {
	validity = CheckDataValidity(year)
	officialCalendarDuringAYear = CheckOfficialRussianCalendarDuringAYear(year)
	if validity == true {
		if officialCalendarDuringAYear == "Julian calendar" {
			isLeapYear = CheckJulianCalendarLeapYear(year)
			if isLeapYear == true {
				dayOfProgrammer = fmt.Sprintf("12.09.%d\n", year)
			} else {
				dayOfProgrammer = fmt.Sprintf("13.09.%d\n", year)
			}
		} else if officialCalendarDuringAYear == "Gregorian calendar" {
			isLeapYear = CheckGregorianCalendarLeapYear(year)
			if isLeapYear == true {
				dayOfProgrammer = fmt.Sprintf("12.09.%d\n", year)
			} else {
				dayOfProgrammer = fmt.Sprintf("13.09.%d\n", year)
			}
		} else {
			isLeapYear = false
			dayOfProgrammer = fmt.Sprintf("26.09.%d\n", year)
		}
	}
	return dayOfProgrammer
}

// Note: fmt.Sprintf formats and returns a string without printing it anywhere.

func main() {
	year = 1918
	dayOfProgrammer = FindDayOfProgrammer(year)
	fmt.Println(dayOfProgrammer)
}
