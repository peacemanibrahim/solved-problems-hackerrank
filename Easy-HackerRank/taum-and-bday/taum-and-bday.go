// Link to problem statement: https://www.hackerrank.com/challenges/taum-and-bday/problem
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	b                                        int32
	w                                        int32
	bc                                       int32
	wc                                       int32
	z                                        int32
	validity                                 bool
	originalCostsOfBlackAndWhiteGifts        int64
	costWhenConvertingBlackGiftsToWhiteGifts int64
	costWhenConvertingWhiteGiftsToBlackGifts int64
	minimumCostOfDikshasGifts                int64
)

func CheckDataValidity(b, w, bc, wc, z int32) bool {
	basePower9 := int32(math.Pow(10, 9))
	if b >= 0 && b <= basePower9 && w >= 0 && w <= basePower9 && bc >= 0 && bc <= basePower9 && wc >= 0 &&
		wc <= basePower9 && z >= 0 && z <= basePower9 {
		validity = true
	} else {
		log.Fatal("Invalid data. check the data.")
	}
	return validity
}

func TaumBday(b, w, bc, wc, z int32) int64 {
	validity = CheckDataValidity(b, w, bc, wc, z)
	if validity == true {
		int64RepOfB := int64(b)
		int64RepOfW := int64(w)
		int64RepOfBC := int64(bc)
		int64RepOfWC := int64(wc)
		int64RepOfZ := int64(z)

		originalCostsOfBlackAndWhiteGifts = (int64RepOfB * int64RepOfBC) + (int64RepOfW * int64RepOfWC)
		costWhenConvertingBlackGiftsToWhiteGifts = (int64RepOfB * int64RepOfBC) + (int64RepOfBC+int64RepOfZ)*int64RepOfW
		costWhenConvertingWhiteGiftsToBlackGifts = (int64RepOfW * int64RepOfWC) + (int64RepOfWC+int64RepOfZ)*int64RepOfB
		minimumCostOfDikshasGifts = originalCostsOfBlackAndWhiteGifts
		if minimumCostOfDikshasGifts > costWhenConvertingBlackGiftsToWhiteGifts {
			minimumCostOfDikshasGifts = costWhenConvertingBlackGiftsToWhiteGifts
		}
		if minimumCostOfDikshasGifts > costWhenConvertingWhiteGiftsToBlackGifts {
			minimumCostOfDikshasGifts = costWhenConvertingWhiteGiftsToBlackGifts
		}
	}
	return minimumCostOfDikshasGifts
}

func main() {
	b = 27984
	w = 1402
	bc = 619246
	wc = 615589
	z = 247954

	minimumCostOfDikshasGifts = TaumBday(b, w, bc, wc, z)
	fmt.Println(minimumCostOfDikshasGifts)
}
