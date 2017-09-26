package main

import "fmt"

func main() {
	var i, num, tmp, sum uint64
	fmt.Scan(&num)

	for i = 0; i < num; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}
	fmt.Println(sum)
}
