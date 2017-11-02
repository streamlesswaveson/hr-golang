package main

import (
	"fmt"
	"unicode"
)

func main() {
	var S string
	fmt.Scan(&S)

	count := 1
	for i := 1; i < len(S); i++ {
		if unicode.IsUpper(rune(S[i])) {
			count++
		}
	}
	fmt.Println(count)
}
