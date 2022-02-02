package main

import (
	"fmt"
	"time"
)

func main() {
	aMPmTimeLayout := "03:04:05PM"
	timeOutputFormat := "15:04:05"
	aMPmTimeInput := "07:05:45PM"

	militaryTime, err := time.Parse(aMPmTimeLayout, aMPmTimeInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(militaryTime)
	fmt.Println(militaryTime.Format(timeOutputFormat))
}
