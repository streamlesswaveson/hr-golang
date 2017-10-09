package main

import "fmt"

func main() {
	var N, high, low, breaksBest, breaksWorst, tmp int

	fmt.Scan(&N)
	fmt.Scan(&high)
	low = high

	for i := 1; i < N; i++ {
		fmt.Scan(&tmp)
		if tmp > high {
			high = tmp
			breaksBest++
		}
		if tmp < low {
			low = tmp
			breaksWorst++
		}
	}

	fmt.Println(breaksBest, breaksWorst)

}
