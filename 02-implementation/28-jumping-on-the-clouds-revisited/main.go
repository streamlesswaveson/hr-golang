package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	clouds := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&clouds[i])
	}

	score := 100
	currentIndex := 0

	for {
		currentIndex = (currentIndex + K) % N
		score -= 1
		if clouds[currentIndex] == 1 {
			score -= 2
		}
		if currentIndex == 0 {
			break
		}
	}
	fmt.Println(score)
}
