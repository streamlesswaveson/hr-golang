package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		N int
	)
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	//fmt.Println(arr)

	if len(arr)%2 == 0 {
		fmt.Println(getMid(arr[0 : len(arr)/2]))
		fmt.Println(getMid(arr))
		fmt.Println(getMid(arr[len(arr)/2:]))
	} else {
		fmt.Println(getMid(arr[0 : len(arr)/2]))
		fmt.Println(getMid(arr))
		fmt.Println(getMid(arr[(len(arr)/2)+1:]))
	}
}

func getMid(arr []int) int {
	if len(arr)%2 == 0 {
		mid := len(arr) / 2
		return (arr[mid] + arr[mid-1]) / 2
	}
	return arr[len(arr)/2]
}
