package main

import "fmt"
import "math/big"

func main() {
	var N int
	fmt.Scan(&N)

	var sum big.Int

	for i := 0; i < N; i++ {
		var x big.Int
		fmt.Scan(&x)
		sum.Add(&sum, &x)
	}

	fmt.Println(sum.String())
}
