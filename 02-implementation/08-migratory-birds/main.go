package main

import "fmt"

func main() {
	var N, tmp, seenMost, seenMostKey int

	fmt.Scan(&N)
	themap := make(map[int]int, 5)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		themap[tmp]++
	}

	for k:=1; k<=5; k++ {
		v := themap[k]
		if v > seenMost {
			seenMost = v
			seenMostKey = k
		}
	}

	fmt.Println(seenMostKey)

}
