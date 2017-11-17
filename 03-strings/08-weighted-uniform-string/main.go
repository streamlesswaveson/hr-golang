package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Scan(&in)

	lookup := map[int32]bool{}
	var start int

	for i:=0; i<len(in); i++ {
		if i == 0 || in[i] != in[i-1]{
			start = i
			lookup[calc([]int32(in[start:start+1]))] = true
			continue
		}
		if in[i] == in[i-1] {
			lookup[calc([]int32(in[start:i+1]))] = true
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

func calc(slice []int32) int32 {
	val := slice[0] -96
	var total int32
	for i:=0; i<len(slice); i++{
		total += val
	}
	return total
}