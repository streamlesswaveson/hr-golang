package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var I, J, K, count int
	fmt.Scan(&I)
	fmt.Scan(&J)
	fmt.Scan(&K)

	for i := I; i <= J; i++ {
		r := reverse(i)
		if int(math.Abs(float64(i)-float64(r)))%K == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func reverse(i int) int {
	s := strconv.Itoa(i)

	rs := ""
	for j := len(s) - 1; j >= 0; j-- {
		rs += string(s[j])
	}
	out, _ := strconv.Atoi(rs)
	return out
}
