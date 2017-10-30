package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		N                     int
		total, mean, variance float64
	)
	fmt.Scan(&N)
	vals := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vals[i])
		total += vals[i]
	}

	mean = total / float64(N)

	total = 0
	for i := 0; i < N; i++ {
		total += math.Pow(vals[i]-mean, 2.0)
	}

	variance = total / float64(N)

	fmt.Printf("%.1f\n", math.Sqrt(variance))

}
