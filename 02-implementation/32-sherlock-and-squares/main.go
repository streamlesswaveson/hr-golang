package main

import "fmt"

func main() {
	var (
		T int
	)
	fmt.Scan(&T)

	mymap := map[int]int{}
	for i := 1; i <= 31622; i++ {
		mymap[i*i] = i
	}

	for i := 0; i < T; i++ {
		solve(mymap)
	}
}

func solve(lookup map[int]int) {
	var A, B, bottom, top int
	fmt.Scan(&A, &B)

	for i:=B; i>=A; i-- {
		if val,ok := lookup[i]; ok {
			top = val
			break
		}
	}
	if top == 0 {
		fmt.Println(0)
		return
	}

	for i:=A-1; i>0; i-- {
		if val,ok := lookup[i]; ok {
			bottom = val
			break
		}
	}

	fmt.Println(top - bottom)


}
