package main

import "fmt"

func main() {
	var n, tmp int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		fmt.Println(grade(tmp))
	}
}

func grade(v int) int {

	if v < 38 {
		return v
	}

	if v%5 == 3 {
		return v + 2
	}

	if v%5 == 4 {
		return v + 1
	}

	return v

}
