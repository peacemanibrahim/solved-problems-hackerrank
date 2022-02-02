// package main

// import (
// 	"fmt"
// 	"log"
// 	"math"
// )

// var (
// 	s        int32
// 	t        int32
// 	a        int32
// 	b        int32
// 	d        int32
// 	validity bool
// )

// func CheckDataValidity(s int32, t int32, a int32, b int32, m []int32, n []int32) bool {
// 	basePower := int32(math.Pow(10, 5))
// 	negativeBasePower := int32(math.Pow(-10, 5))
// 	numberOfApples := int32(len(m))
// 	numberOfOranges := int32(len(n))
// 	if s >= 1 && s <= basePower && t >= 1 && t <= basePower && a >= 1 && a <= basePower && b >= 1 && b <= basePower &&
// 		numberOfApples >= 1 && numberOfApples <= basePower && numberOfOranges >= 1 && numberOfOranges <= basePower &&
// 		s > a && t > s && b > t {
// 		for i := range m {
// 			if m[i] >= negativeBasePower && m[i] <= basePower {

// 			} else {
// 				log.Fatal("Invalid data for apples array")
// 			}
// 		}

// 		for i := range n {
// 			if n[i] >= negativeBasePower && n[i] <= basePower {

// 			} else {
// 				log.Fatal("Invalid data for the oranges array")
// 			}
// 		}
// 		validity = true
// 	} else {
// 		log.Fatal("Invalid data. Check the data")
// 	}
// 	return validity
// }

// func main() {
// 	s = 7
// 	t = 10
// 	a = 4
// 	b = 12
// 	m := []int32{2, 3, -4}
// 	n := []int32{3, -2, -4}
// 	var applesLandingPositionList []int32
// 	var orangesLandingPositionList []int32
// 	var applesWithinSamHouse []int32
// 	var orangesWithinSamHouse []int32
// 	validity = CheckDataValidity(s, t, a, b, m, n)

// 	if validity == true {
// 		for i := range m {
// 			positionAppleLandedAt := m[i] + a
// 			applesLandingPositionList = append(applesLandingPositionList, positionAppleLandedAt)
// 		}

// 		for i := range applesLandingPositionList {
// 			if applesLandingPositionList[i] >= s && applesLandingPositionList[i] <= t {
// 				applesWithinSamHouse = append(applesWithinSamHouse, applesLandingPositionList[i])
// 			}
// 		}

// 		for i := range n {
// 			positionOrangeLandedAt := n[i] + b
// 			orangesLandingPositionList = append(orangesLandingPositionList, positionOrangeLandedAt)
// 		}

// 		for i := range orangesLandingPositionList {
// 			if orangesLandingPositionList[i] >= s && orangesLandingPositionList[i] <= t {
// 				orangesWithinSamHouse = append(orangesWithinSamHouse, orangesLandingPositionList[i])
// 			}
// 		}

// 		numberOfApplesWithinSamHouse := len(applesWithinSamHouse)
// 		numberOfOrangesWithinSamHouse := len(orangesWithinSamHouse)

// 		fmt.Println(numberOfApplesWithinSamHouse)
// 		fmt.Println(numberOfOrangesWithinSamHouse)

// 		// fmt.Println(applesWithinSamHouse)
// 		// fmt.Println(orangesWithinSamHouse)

// 		// fmt.Println(applesLandingPositionList)
// 		// fmt.Println(orangesLandingPositionList)
// 	}

// }

// // Definition of terms: s and t is the position of Sam's house. The apple tree is located at a and the orange tree at
// // b. m is a list of the unit distances of the apple from the apple Tree. m is a list of the distances of the orange from
// // the orange Tree.
// // Note that each distance in m is the unit distance of apple from a, it is not the position the
// // apple landed at. To get the position the apple landed at you add the unit distance of the apple from a and the
// // position of the tree
