package main

import "fmt"

func main() {
	var N, tmp int
	fmt.Scan(&N)

	inverseP := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&tmp)
		inverseP[tmp] = i
	}
	fmt.Println(inverseP)

	for i := 1; i <= N; i++ {
		fmt.Println(inverseP[inverseP[i]])
	}
}
