package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i:=0; i<N; i++ {
		fmt.Println(solve())
	}
}

func solve() string {
	var S string
	fmt.Scan(&S)

	j := len(S)-1
	for i:=0; i<len(S)-1;i++{
		if math.Abs(float64(S[i]) - float64(S[i+1])) != math.Abs(float64(S[j]) - float64(S[j-1])) {
			return "Not Funny"
		}
		j--
	}
	return "Funny"

}