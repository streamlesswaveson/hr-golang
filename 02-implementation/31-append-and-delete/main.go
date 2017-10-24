package main

import "fmt"

func main() {
	var (
		S, T string
		K    int
	)
	fmt.Scan(&S, &T, &K)

	i := 0
	for i = 0; i < len(S) && i < len(T); i++ {
		if S[i] != T[i] {
			break
		}
	}
	if len(S[i:])+len(T[i:]) == K {
		fmt.Println("Yes",1)
	} else if ((len(S)+len(T)+1)-K)%2 == 0 {
		fmt.Println("Yes",2)
	} else if S == T && K%2 == 0 {
		fmt.Println("Yes",3)
	} else {
		fmt.Println("No")
	}
}
