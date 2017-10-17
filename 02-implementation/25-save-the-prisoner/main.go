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
	var N, M, S int
	fmt.Scan(&N)
	fmt.Scan(&M)
	fmt.Scan(&S)

	last := M % N

	result := (last + (S - 1))%N
	if result == 0 {
		fmt.Println(N)
	} else {
		fmt.Println(result)
	}

}
