package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)
	found := -1
	for i := 0; i < N-1; i++ {
		if arr[i] == arr[i+1] {
			i++
			continue
		}
		found = arr[i]
		break
	}
	if found < 0 {
		fmt.Println(arr[len(arr)-1])
	} else {
		fmt.Println(found)
	}
}
