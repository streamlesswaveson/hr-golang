package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)

	for i := 1; i <= size; i++ {
		spaces := size - i
		hashes := i
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < hashes; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
