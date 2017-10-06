package main

import (
	"fmt"
	"math"
)

func main() {
	var k1, v1, k2, v2 int
	fmt.Scan(&k1)
	fmt.Scan(&v1)
	fmt.Scan(&k2)
	fmt.Scan(&v2)

	getAbsDistance := func(k1 int, k2 int) int {
		dist := math.Abs(float64(k2) - float64(k1))
		return int(dist)

	}

	dist := getAbsDistance(k1, k2)
	for {
		k1 += v1
		k2 += v2
		if k1 == k2 {
			fmt.Println("YES")
			break
		}
		tmp := getAbsDistance(k1, k2)
		// if distance is greater than last iteration
		if tmp > dist {
			fmt.Println("NO")
			break
		}
		// if distance remains constant
		if tmp == dist {
			fmt.Println("NO")
			break
		}
		dist = tmp

	}
}
