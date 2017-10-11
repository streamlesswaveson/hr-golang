package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Println(processQuery())
	}
}

func processQuery() string {
	var A, B, M int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&M)

	bDist := math.Abs(float64(M - B))
	aDist := math.Abs(float64(M - A))
	if aDist == bDist {
		return "Mouse C"
	}
	if aDist < bDist {
		return "Cat A"
	}
	return "Cat B"
}
