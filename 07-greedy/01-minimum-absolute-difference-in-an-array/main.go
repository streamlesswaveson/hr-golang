package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Float64s(arr)

	least := math.MaxFloat64

	j := 1
	for i := 0; i < N-1; i++ {
		temp := math.Abs(arr[i] - arr[j])
		if temp < least {
			least = temp
		}
		j++
	}
	fmt.Println(least)
}
