// Link to problem statement: https://www.hackerrank.com/challenges/cats-and-a-mouse/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	x                int32
	y                int32
	z                int32
	validity         bool
	catToArriveFirst string
)

func CheckDataValidity(x, y, z int32) bool {
	if x >= 1 && x <= 100 && y >= 1 && y <= 100 && z >= 1 && z <= 100 {
		validity = true
	} else {
		log.Fatal("Invalid data. Check the data.")
	}
	return validity
}

func CatAndMouse(x, y, z int32) string {
	validity = CheckDataValidity(x, y, z)
	catAPosition := x
	catBPosition := y
	mouseCPosition := z
	catADistanceFromMouseC := math.Abs(float64(mouseCPosition - catAPosition))
	catBDistanceFromMouseC := math.Abs(float64(mouseCPosition - catBPosition))
	if catADistanceFromMouseC < catBDistanceFromMouseC {
		catToArriveFirst = "Cat A"
	} else if catBDistanceFromMouseC < catADistanceFromMouseC {
		catToArriveFirst = "Cat B"
	} else if catADistanceFromMouseC == catBDistanceFromMouseC {
		catToArriveFirst = "Mouse C"
	}
	return catToArriveFirst
}

func main() {
	x = 1
	y = 3
	z = 2

	catToArriveFirst = CatAndMouse(x, y, z)
	fmt.Println(catToArriveFirst)
}
