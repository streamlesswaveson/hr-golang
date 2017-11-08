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
	var S string
	fmt.Scan(&S)

	ref := "hackerrank"
	index := 0
	for _, s := range S {
		if index >= len(ref) {
			break
		}
		if s == rune(ref[index]) {
			index++
		}
	}
	if index >= len(ref) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
