// Link to problem statement: https://www.hackerrank.com/challenges/time-conversion/problem
package main

import (
	"fmt"
	"log"
	"time"
)

func timeConversion(aMPmTimeInput string) string {
	aMPmTimeLayout := "03:04:05PM"
	timeOutputFormat := "15:04:05"

	militaryTime, err := time.Parse(aMPmTimeLayout, aMPmTimeInput)
	if err != nil {
		log.Fatal(err)
	}
	return militaryTime.Format(timeOutputFormat)
}

func main() {
	result := timeConversion("07:05:45PM")
	fmt.Println(result)
}
