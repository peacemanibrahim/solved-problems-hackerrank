// Link to problem statement: https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	validity                        bool
	minimumScore                    int32
	maximumScore                    int32
	minimumScoreCount               int32
	maximumScoreCount               int32
	minimumAndMaximumScoreCountList []int32
)

func CheckDataValidity(scores []int32) bool {
	scoresLength := len(scores)
	basePower := int32(math.Pow(10, 8))
	if scoresLength >= 1 && scoresLength <= 1000 {
		for i := range scores {
			if scores[i] >= 0 && scores[i] <= basePower {
				validity = true
			} else {
				log.Fatal("Invalid scores data")
			}
		}
	} else {
		log.Fatal("Invalid number of games")
	}
	return validity
}

func BreakingRecords(scores []int32) []int32 {
	validity = CheckDataValidity(scores)
	minimumScore = scores[0]
	maximumScore = scores[0]
	minimumScoreCount = 0
	maximumScoreCount = 0
	scoresLength := len(scores)

	if validity == true {
		for i := 1; i <= scoresLength-1; i++ {
			if scores[i] > maximumScore {
				maximumScore = scores[i]
				maximumScoreCount++
			} else if scores[i] < minimumScore {
				minimumScore = scores[i]
				minimumScoreCount++
			}
		}
		minimumAndMaximumScoreCountList = append(minimumAndMaximumScoreCountList, maximumScoreCount)
		minimumAndMaximumScoreCountList = append(minimumAndMaximumScoreCountList, minimumScoreCount)
	}
	return minimumAndMaximumScoreCountList
}

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	minimumAndMaximumScoreCountList = BreakingRecords(scores)

	fmt.Println(minimumAndMaximumScoreCountList)
}
