package main

import "fmt"

func main() {
	var N, valleys, alt int
	var T string

	fmt.Scan(&N)
	fmt.Scan(&T)

	//fmt.Println(T);
	for i := 0; i < N; i++ {
		if string(T[i]) == "U" {
			alt++
			if alt == 0 {
				valleys++
			}
		} else if string(T[i]) == "D" {
			alt--
		}
	}
	fmt.Println(valleys)
}
