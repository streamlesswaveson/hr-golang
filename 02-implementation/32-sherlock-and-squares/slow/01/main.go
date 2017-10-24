package main

import "fmt"

func main() {
	var (
		T int
	)
	fmt.Scan(&T)

	mymap := map[int]bool{}
	for i := 1; i <= 31622; i++ {
		mymap[i*i] = true
	}

	for i := 0; i < T; i++ {
		solve(mymap)
	}
}

func solve(lookup map[int]bool) {
	var A, B int
	fmt.Scan(&A, &B)

	var count int
	for i:=A; i<=B; i++ {
		if lookup[i] {
			count++
		}
	}

	fmt.Println(count)

}
