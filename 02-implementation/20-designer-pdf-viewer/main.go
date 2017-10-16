package main

import "fmt"

func main() {
	var (
		tmp, tallest int
		line         string
	)
	abc := map[rune]int{}

	for i := 'a'; i <= 'z'; i++ {
		fmt.Scan(&tmp)
		abc[i] = tmp
	}
	fmt.Scan(&line)

	for _, t := range line {
		if abc[t] > tallest {
			tallest = abc[t]
		}
	}
	fmt.Println(tallest * len(line))
}
