package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	count := 0
	for i := 0; i < len(S); i = i + 3 {
		s := S[i : i+3]
		if s == "SOS" {
			continue
		}
		if s[0] != 'S' {
			count++
		}
		if s[1] != 'O' {
			count++
		}
		if s[2] != 'S' {
			count++
		}
	}
	fmt.Println(count)
}
