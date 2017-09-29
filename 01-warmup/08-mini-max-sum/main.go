package main

import (
	"fmt"
	"sort"
)

type mytype []uint64

func (m mytype) Len() int {
	return len(m)
}

func (m mytype) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m mytype) Less(i, j int) bool {
	return m[i] < m[j]
}

func main() {
	arr := make(mytype, 5)

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Sort(arr)

	var sumA, sumB uint64
	for i := 0; i < 4; i++ {
		sumA += arr[i]
	}
	for i := 1; i < 5; i++ {
		sumB += arr[i]
	}
	fmt.Println(sumA, sumB)
}
