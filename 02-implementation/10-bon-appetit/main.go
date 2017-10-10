package main

import "fmt"

func main() {
	var N, K, tmp, actual, charged int
	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		if i != K {
			actual += tmp
		}
	}
	fmt.Scan(&charged)

	if charged == actual/2 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged - (actual / 2))
	}

}
