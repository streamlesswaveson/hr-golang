package main

import (
	"fmt"
	"math"
)

func main() {
	arraySize := 0
	total1 := 0
	total2 := 0

	fmt.Scan(&arraySize)
	numLines := arraySize

	{
		var d int
		for j := 0; j < numLines; j++ {
			for i := 0; i < arraySize; i++ {
				_, err := fmt.Scan(&d)
				if err != nil {
					break
				}
				if i == j {
					total1 += d
				}
				if i+1 == numLines-j {
					total2 += d
				}

			}
		}
	}
	result := total1 - total2
	fmt.Println(int(math.Abs(float64(result))))

}
