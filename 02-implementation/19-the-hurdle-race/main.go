package main

import "fmt"

func main() {
	var N, K, tmp, max int
	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		if tmp > K {
			d := tmp - K
			if d > max {
				max = d
			}
		}
	}
	fmt.Println(max)

}
