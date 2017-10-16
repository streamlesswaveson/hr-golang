package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		solve()
	}
}

func solve() {
	var N, K, tmp, count int
	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		if tmp <= 0 {
			count++
		}
	}
	if count >= K {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
