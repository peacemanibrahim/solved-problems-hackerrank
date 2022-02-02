// Link to problem statement: https://www.hackerrank.com/challenges/strange-code/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	t                                                         int64
	validity                                                  bool
	counterTimeInSecondsAtTheEndOfEachCycle                   int64
	counterDisplayValueAtTheBeginningOfEachCycle              int64
	counterDisplayValueAtTheBeginningOfPreviousCycle          int64
	valueDisplayedByTheCounterAtTimeT                         int64
	differenceBtwCounterTimeInSecsAtTheEndOfEachCycleAndTimeT int64
)

const (
	valueDisplayedByTheCounterAtTheEndofEachCycle = 1
)

func CheckDataValidity(t int64) bool {
	basePower12 := int64(math.Pow(10, 12))
	if t >= 1 && t <= basePower12 {
		validity = true
	} else {
		log.Fatal("Invalid data for t. Check the data.")
	}
	return validity
}

func StrangeCounter(t int64) int64 {
	counterDisplayValueAtTheBeginningOfPreviousCycle = 3
	counterTimeInSecondsAtTheEndOfEachCycle = 3
	validity = CheckDataValidity(t)
	if validity == true {
		for counterTimeInSecondsAtTheEndOfEachCycle < t {
			counterDisplayValueAtTheBeginningOfEachCycle = counterDisplayValueAtTheBeginningOfPreviousCycle * 2
			counterTimeInSecondsAtTheEndOfEachCycle += counterDisplayValueAtTheBeginningOfEachCycle
			counterDisplayValueAtTheBeginningOfPreviousCycle = counterDisplayValueAtTheBeginningOfEachCycle
		}
		if t == counterTimeInSecondsAtTheEndOfEachCycle {
			return valueDisplayedByTheCounterAtTheEndofEachCycle
		} else if t < counterTimeInSecondsAtTheEndOfEachCycle {
			differenceBtwCounterTimeInSecsAtTheEndOfEachCycleAndTimeT = counterTimeInSecondsAtTheEndOfEachCycle - t
			valueDisplayedByTheCounterAtTimeT = differenceBtwCounterTimeInSecsAtTheEndOfEachCycleAndTimeT +
				valueDisplayedByTheCounterAtTheEndofEachCycle
		}
	}
	return valueDisplayedByTheCounterAtTimeT
}

func main() {
	t = 21

	valueDisplayedByTheCounterAtTimeT = StrangeCounter(t)
	fmt.Println(valueDisplayedByTheCounterAtTimeT)
}

// Explanation: Since I know that when "t" is equal to the counter time in seconds at the end of each cycle, the
// display value is equal to 1. All I need to do is check when "t" is less than the counter time in seconds at the end
// of each cycle and by how much is "t" less.
// For example: when t = 9 (9 is the counter time in seconds at the end of the 2nd cycle), display value is
// equal to 1.
// So when t = 7 (7 is less than 9), so I'll find the difference between t and counter time in seconds at
// the end of each cycle, which is 9 - 7 = 2. Therefore the display value of t = 7 is the addition of the
// "difference" and the display value at the end of each cycle which is 1. 2 + 1 = 3
