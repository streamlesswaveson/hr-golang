package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		N, F int
		arr  []float64
	)
	fmt.Scan(&N)
	tmp := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&F)
		for j := 0; j < F; j++ {
			arr = append(arr, tmp[i])
		}
	}
	sort.Float64s(arr)
	//fmt.Println(arr)

	if len(arr)%2 == 0 {
		q1Mid := getMid(arr[0 : len(arr)/2])
		q3Mid := getMid(arr[len(arr)/2:])
		print(q3Mid - q1Mid)
	} else {
		q1Mid := getMid(arr[0 : len(arr)/2])
		q3Mid := getMid(arr[len(arr)/2+1:])
		print(q3Mid - q1Mid)
	}
}

func print(v float64) {
	fmt.Printf("%.1f\n", v)
}

func getMid(arr []float64) float64 {
	if len(arr)%2 == 0 {
		mid := len(arr) / 2
		return (arr[mid] + arr[mid-1]) / 2
	}
	return arr[len(arr)/2]
}
