package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Scan(&in)

	lookup := map[int32]bool{}
	var start int
	var lastsum int32

	for i:=0; i<len(in); i++ {
		if i == 0 || in[i] != in[i-1]{
			start = i
			lastsum = int32(in[start]) - 96
			lookup[lastsum] = true
			continue
		}
		if in[i] == in[i-1] {
			lastsum = lastsum + (int32(in[start]) - 96)
			lookup[lastsum] = true
			continue
		}
	}
	var N int
	fmt.Scan(&N)

	for i:=0; i<N; i++ {
		solve(lookup)
	}
}

func solve(lookup map[int32]bool) {
	var tmp int32
	fmt.Scan(&tmp)
	if lookup[tmp] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
