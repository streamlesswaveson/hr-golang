package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	ff := p / 2
	fb := ((n + 1 - n%2) - p) / 2
	if ff < fb {
		fmt.Println(ff)
	} else {
		fmt.Println(fb)
	}
}