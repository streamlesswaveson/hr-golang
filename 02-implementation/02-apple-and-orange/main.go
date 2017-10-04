package main

import (
	"fmt"
)

func main() {
	var samLeft, samRight, appleTree, orangeTree, totalApples, totalOranges int

	fmt.Scan(&samLeft)
	fmt.Scan(&samRight)
	fmt.Scan(&appleTree)
	fmt.Scan(&orangeTree)
	fmt.Scan(&totalApples)
	fmt.Scan(&totalOranges)

	appleOut := eval(totalApples, appleTree, samLeft, samRight)
	orangeOut := eval(totalOranges, orangeTree, samLeft, samRight)

	fmt.Println(len(appleOut))
	fmt.Println(len(orangeOut))
}

func eval(total int, tree int, left int, right int) []int {
	var tmp int

	var out = []int{}
	for i := 0; i < total; i++ {
		fmt.Scan(&tmp)
		var d = tree + tmp
		if left <= d && d <= right {
			out = append(out, d)
		}
	}

	return out
}
