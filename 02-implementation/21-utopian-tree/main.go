package main

import "fmt"

func main() {
	var N, cycles int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&cycles)
		if cycles == 0 {
			fmt.Println(1)
			continue
		}

		spring := true
		height := 1
		for j := 0; j < cycles; j++ {

			if spring {
				height *= 2
				spring = false
			} else {
				height += 1
				spring = true
			}
		}
		fmt.Println(height)
	}
}
