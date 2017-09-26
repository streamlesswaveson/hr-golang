package main

import "fmt"

func main() {
	alice := prepare()
	bob := prepare()

	var aliceTotal, bobTotal uint64
	for i := 0; i < 3; i++ {
		if alice[i] == bob[i] {
			continue
		}
		if alice[i] > bob[i] {
			aliceTotal++
			continue
		}
		if bob[i] > alice[i] {
			bobTotal++
			continue
		}
	}
	fmt.Printf("%d %d\n", aliceTotal, bobTotal)
}

func prepare() []uint64 {
	var n uint64
	arr := make([]uint64, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&n)
		arr[i] = n
	}
	return arr
}
