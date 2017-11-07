package main

import "fmt"

func main() {
	var (
		N, K int32
		S    string
	)
	fmt.Scan(&N, &S, &K)

	K = K % 26
	var out string
	for _, s := range S {
		var letter rune
		if 'a' <= s && s <= 'z' {
			if s+K > 'z' {
				letter = 'a' + ((s + K) - 'z') - 1
				out += string(letter)
			} else {
				out += string(s + K)
			}

		} else if 'A' <= s && s <= 'Z' {
			if s+K > 'Z' {
				letter = 'A' + ((s + K) - 'Z') - 1
				out += string(letter)
			} else {
				out += string(s + K)
			}

		} else {
			out += string(s)
		}
	}
	fmt.Println(out)
}
