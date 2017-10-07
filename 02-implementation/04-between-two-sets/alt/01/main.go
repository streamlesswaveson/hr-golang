package main

import "fmt"

func areFactors(i int, a []int) bool {
	for _, elem := range a {
		if elem % i != 0 {
			return false
		}
	}
	return true
}

func isFactor(i int, a []int) bool {
	for _, elem := range a {
		if i % elem != 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, m, tmp, i, count int
	maxA := 1
	minB := 100
	fmt.Scanf("%d %d", &n, &m)
	A := make([]int, n)
	B := make([]int, m)
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		A[i] = tmp
		if tmp > maxA {
			maxA = tmp
		}
	}
	for i = 0; i < m; i++ {
		fmt.Scanf("%d", &tmp)
		B[i] = tmp
		if tmp < minB {
			minB = tmp
		}
	}

	for i = maxA; i <= minB; i++ {
		if isFactor(i, A) && areFactors(i, B) {
			count++
		}
	}
	fmt.Println(count)
}