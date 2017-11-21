package main

import "fmt"

// https://www.hackerrank.com/challenges/lonely-integer/editorial
func main()  {

	var N,tmp int
	fmt.Scan(&N)

	result :=0
	for i:=0; i<N; i++ {
		fmt.Scan(&tmp)
		result ^= tmp
	}
	fmt.Println(result)

}