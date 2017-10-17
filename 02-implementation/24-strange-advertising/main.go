package main

import "fmt"

func main() {
	var N, people, sum int
	fmt.Scan(&N)

	people = 5
	sum = 0
	for i := 0; i < N; i++ {
		people = (people / 2)
		sum += people
		people *= 3
	}

	fmt.Println(sum)

}
