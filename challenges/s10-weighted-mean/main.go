package main

import "fmt"

func main() {
	var (
		N                       int
		totalWeight, total, tmp float64
	)
	fmt.Scan(&N)

	values := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&values[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		totalWeight += tmp
		total += (values[i] * tmp)
	}
	fmt.Printf("%.1f\n", total / totalWeight)
}
