package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	var N int

	fmt.Scan(&N)

	cupcakes := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&cupcakes[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cupcakes)))

	var miles float64
	for i, v := range  cupcakes {
		miles = miles + (float64(v) * math.Pow(2, float64(i)))
	}
	fmt.Println(int64(miles))
}