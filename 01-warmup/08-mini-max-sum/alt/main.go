package main

import (
	"fmt"
	"math"
)

func main() {
	var s = 0
	var min, max = math.MaxInt32, math.MinInt32
	for i := 0; i < 5; i++ {
		var x int
		fmt.Scan(&x)

		if x < min {
			min = x
		}

		if x > max {
			max = x
		}

		s += x
	}

	fmt.Println(s-max, s-min)
}