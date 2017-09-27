package main

import (
	"fmt"
)

func main() {
	var i, size, tmp, negTotal, posTotal, zeroTotal float64

	fmt.Scan(&size)

	for i = 0; i < size; i++ {
		fmt.Scan(&tmp)
		if tmp > 0 {
			posTotal++
		} else if tmp < 0 {
			negTotal++
		} else {
			zeroTotal++
		}
	}

	fmt.Printf("%f\n", posTotal / size)
	fmt.Printf("%f\n", negTotal / size)
	fmt.Printf("%f\n", zeroTotal / size)
}
