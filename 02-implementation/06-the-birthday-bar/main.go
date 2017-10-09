package main

import "fmt"

func main() {
	var N, D, M, count int

	fmt.Scan(&N)

	bar := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&bar[i])
	}

	fmt.Scan(&D)
	fmt.Scan(&M)

	for i := 0; i < N; i++ {
		if i+M > N {
			break
		}
		s := bar[i : i+M]
		sum := sum(s)
		if sum == D {
			count++
		}
	}

	fmt.Println(count)

}

func sum(a []int) int {
	var sum int
	for _, a := range a {
		sum += a
	}
	return sum
}
