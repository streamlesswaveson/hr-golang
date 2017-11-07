package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	fmt.Scanln(&k)

	k %= 26
	output := make([]rune, n)
	for i, c := range s {
		output[i] = c
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			output[i] += rune(k)
			if 'a' <= s[i] && s[i] <= 'z' && output[i] > 'z' {
				output[i] -= ('z' - 'a' + 1)
			} else if 'A' <= s[i] && s[i] <= 'Z' && output[i] > 'Z' {
				output[i] -= ('Z' - 'A' + 1)
			}
		}
	}
	fmt.Println(string(output))
}