package main

import (
	"fmt"
	"sort"
)

// example
// 1 1
// 1
// 100

// answer = 9

func main() {

	var sizeA, sizeB int
	fmt.Scan(&sizeA)
	fmt.Scan(&sizeB)

	A := scanArr(sizeA)
	B := scanArr(sizeB)

	cands := factorize(B[0])
	success := []int{}

	//fmt.Println(cands)

	for i := 0; i < len(cands); i++ {
		if isDivisble(cands[i], A) && canDivide(cands[i], B) {
			success = append(success, cands[i])
		}
	}
	//fmt.Println(success)

	fmt.Println(len(success))
}

func factorize(a int) []int {
	out := []int{}
	for i := a; i > 0; i-- {
		if a%i == 0 {
			out = append(out, i)
		}
	}
	return out
}

func isDivisble(cand int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if cand%arr[i] != 0 {
			return false
		}
	}
	return true
}

func canDivide(cand int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%cand != 0 {
			return false
		}
	}
	return true
}

func scanArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	return arr
}
