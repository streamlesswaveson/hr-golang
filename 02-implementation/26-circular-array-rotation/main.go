package main

import "fmt"

func main() {
	var N, K, Q, tmp int
	fmt.Scan(&N, &K, &Q)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[(i+K)%N])
	}

	for i := 0; i < Q; i++ {
		fmt.Scan(&tmp)
		fmt.Println(arr[tmp])
	}

}
