package main

import "fmt"

func main() {
	var N, tmp int
	fmt.Scan(&N)

	m := map[int]int{}

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		m[tmp]++
	}

	max := 0
	for i := 0; i < 100; i++ {
		sum := 0
		if _, ok := m[i]; ok  {
			sum += m[i]
		} else {
			continue
		}
		if _,ok := m[i+1]; ok && i !=99 {
			sum += m[i+1]
		}

		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
