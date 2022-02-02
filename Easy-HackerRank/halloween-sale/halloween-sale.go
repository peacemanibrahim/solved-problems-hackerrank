// Link to problem statement: https://www.hackerrank.com/challenges/halloween-sale/problem
package main

import (
	"fmt"
	"log"
)

var (
	p                             int32
	d                             int32
	m                             int32
	s                             int32
	validity                      bool
	sumOfCostOfVideoGames         int32
	noOfVideoGamesThatCanBeBought int32
)

func CheckDataValidity(p, d, m, s int32) bool {
	if m >= 1 && p >= m && p <= 100 && d >= 1 && d <= 100 && s >= 1 && s <= 10000 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func HowManyGames(p, d, m, s int32) int32 {
	if s < p {
		noOfVideoGamesThatCanBeBought = 0
		return noOfVideoGamesThatCanBeBought
	}
	noOfVideoGamesThatCanBeBought = 1
	validity = CheckDataValidity(p, d, m, s)
	if validity == true {
		sumOfCostOfVideoGames = p
		for sumOfCostOfVideoGames <= s {
			if p-d > m {
				p = p - d
			} else if p-d <= m {
				p = m
			}
			sumOfCostOfVideoGames += p
			if sumOfCostOfVideoGames <= s {
				noOfVideoGamesThatCanBeBought++
			}
		}
	}
	return noOfVideoGamesThatCanBeBought
}

func main() {
	p = 100
	d = 1
	m = 1
	s = 99

	noOfVideoGamesThatCanBeBought = HowManyGames(p, d, m, s)
	fmt.Println(noOfVideoGamesThatCanBeBought)
}

// Note: p is the price of the first video game.
// Note: d is the discount from the previous video game price.
// Note: m is the minimum cost of a video game.
// Note: s is the starting budget
