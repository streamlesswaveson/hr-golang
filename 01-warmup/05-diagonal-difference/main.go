package main

import (
	"fmt"
	"math"
)

func main() {
	var size, aIndex, bIndex int
	var tmp, aSum, bSum float64
	fmt.Scan(&size)
	aIndex = 0
	bIndex = size - 1

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&tmp)
			if j == aIndex {
				aSum += tmp
			}
			if j == bIndex {
				bSum += tmp
			}
		}
		aIndex++
		bIndex--
	}
	// This is incorrect - do not take the individual absolute values of each sum beforehand
	//fmt.Println(math.Abs(math.Abs(aSum) - math.Abs(bSum)))
	fmt.Println(math.Abs(aSum - bSum))
}
