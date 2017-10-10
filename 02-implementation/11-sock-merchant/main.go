package main

import "fmt"

func main() {
	var N, tmp, howmany int
	fmt.Scan(&N)

	mymap := make(map[int]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		mymap[tmp]++
		if mymap[tmp]%2 == 0 {
			howmany++
		}
	}
	fmt.Println(howmany)
}
