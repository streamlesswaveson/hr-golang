package main

import (
	"fmt"
	"sort"
)

func main() {
	var money, N, M int

	fmt.Scan(&money)
	fmt.Scan(&N)
	fmt.Scan(&M)

	nPrices := scanPrices(N)
	mPrices := scanPrices(M)

	sort.Ints(nPrices)
	sort.Ints(mPrices)

	total := -1
	for i := 0; i < N; i++ {
		k := nPrices[i]
		if k >= money {
			break
		}
		for j := 0; j < M; j++ {
			u := mPrices[j]
			if u >= money {
				break
			}
			c := k + u
			if c > total && c <= money {
				total = c
			}
		}

	}

	fmt.Println(total)

}

func scanPrices(howmany int) []int {
	p := make([]int, howmany)
	for i := 0; i < howmany; i++ {
		fmt.Scan(&p[i])
	}
	return p
}
