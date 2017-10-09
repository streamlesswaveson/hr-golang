package main

import "fmt"

func main() {
	var N, K, count int
	fmt.Scan(&N)
	fmt.Scan(&K)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if (arr[i]+arr[j])%K == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
