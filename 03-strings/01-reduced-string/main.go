package main

import "fmt"

func main() {
	var (
		S string
	)
	fmt.Scan(&S)

	lastLength := len(S)
	for {
		S = reduce(S)
		if lastLength == len(S) {
			break
		}
		lastLength = len(S)
	}
	if S == "" {
		fmt.Println("Empty String")
	} else {
		fmt.Println(S)
	}
}

func reduce(S string) string {
	var out string
	for i := 0; i < len(S); i++ {
		if i+1 < len(S) && S[i] == S[i+1] {
			i++
			continue
		}
		out += string(S[i])
	}
	return out
}
