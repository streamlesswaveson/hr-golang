package main

import (
	"fmt"
	"math"
)

func main() {
	var n, count, tmp int
	fmt.Scan(&n)

	currentMax := math.MinInt32

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp > currentMax {
			currentMax = tmp
			count = 1
		} else if currentMax == tmp {
			count += 1
		}
	}
	fmt.Println(count)
}
