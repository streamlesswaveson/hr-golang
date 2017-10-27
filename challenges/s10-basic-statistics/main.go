package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		N, maxSeen      int
		sum, currentMin float64
	)
	fmt.Scan(&N)

	arr := make([]float64, N)
	currentMin = math.MaxFloat64
	seen := map[float64]int{}
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
		seen[arr[i]]++
		sum += arr[i]

		if seen[arr[i]] > maxSeen {
			maxSeen = seen[arr[i]]
			currentMin = arr[i]
		} else if seen[arr[i]] == maxSeen {
			if arr[i] < currentMin {
				maxSeen = seen[arr[i]]
				currentMin = arr[i]
			}
		}
	}
	sort.Float64s(arr)

	fmt.Println(sum / float64(N))
	if N%2 == 0 {
		half := N / 2
		fmt.Println((arr[half] + arr[half-1]) / 2)
	} else {
		fmt.Print(arr[N/2])
	}
	fmt.Println(currentMin)

}
