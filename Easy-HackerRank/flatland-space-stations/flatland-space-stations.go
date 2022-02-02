// Link to problem statement: https://www.hackerrank.com/challenges/flatland-space-stations/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	n                                               int32
	validity                                        bool
	nearestDistanceBtwCityAndSpaceStation           int32
	maximumDistanceBtwAnyCityAndNearestSpaceStation int32
)

func CheckDataValidity(n int32, c []int32) bool {
	m := int32(len(c))
	if n >= 1 && n <= 100000 && m >= 1 && m <= n {
		validity = true
	} else {
		log.Fatal("Invalid data for n and/or m")
	}
	return validity
}

func FlatlandSpaceStations(n int32, c []int32) int32 {
	nearestDistanceBtwCityAndSpaceStation = 99999
	maximumDistanceBtwAnyCityAndNearestSpaceStation = 0
	validity = CheckDataValidity(n, c)
	if validity == true {
		for i := 0; i < int(n); i++ {
			for j := 0; j < len(c); j++ {
				distanceBtwCityAndSpaceStation := math.Abs(float64(i - int(c[j])))
				if int32(distanceBtwCityAndSpaceStation) < nearestDistanceBtwCityAndSpaceStation {
					nearestDistanceBtwCityAndSpaceStation = int32(distanceBtwCityAndSpaceStation)
				}
			}
			if nearestDistanceBtwCityAndSpaceStation > maximumDistanceBtwAnyCityAndNearestSpaceStation {
				maximumDistanceBtwAnyCityAndNearestSpaceStation = nearestDistanceBtwCityAndSpaceStation
			}
			nearestDistanceBtwCityAndSpaceStation = 99999
		}
	}
	return maximumDistanceBtwAnyCityAndNearestSpaceStation
}

func main() {
	n = 5
	c := []int32{0, 4}

	maximumDistanceBtwAnyCityAndNearestSpaceStation = FlatlandSpaceStations(n, c)
	fmt.Println(maximumDistanceBtwAnyCityAndNearestSpaceStation)
}
