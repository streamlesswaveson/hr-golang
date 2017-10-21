package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		solve()
	}
}

func solve() {
	var C int
	fmt.Scan(&C)

	s := strconv.Itoa(C)

	ticks :=0
	for _, s := range s {
		i, _ := strconv.Atoi(string(s))

		if i !=0  && C%i==0 {
			ticks++
		}
	}
	fmt.Println(ticks)
}
