package main

import (
	"fmt"
)

func main() {
	var H, M, S int
	var suffix string
	fmt.Scanf("%d:%d:%d%s", &H, &M, &S, &suffix)
	if suffix == "PM" && H < 12 {
		H += 12
	} else if suffix == "AM" && H >= 12 {
		H -= 12
	}
	fmt.Printf("%02d:%02d:%02d\n", H%24, M, S)
}
