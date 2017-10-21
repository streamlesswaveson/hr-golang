package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int64
	fmt.Scan(&N)

	fmt.Println(fact(big.NewInt(N)))
}

// https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
// and https://play.golang.org/p/feacvk4P4O
func fact(y *big.Int) (result *big.Int) {
	if y.Cmp(big.NewInt(0)) == 0 || y.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}
	result = new(big.Int)
	result.Set(y)
	result.Mul(result, fact(y.Sub(y, big.NewInt(1))))
	return result
}
