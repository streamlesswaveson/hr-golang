package main

import "fmt"

func main() {
	var year int
	fmt.Scan(&year)

	isGregorian := true
	isLeap := false

	if year < 1918 {
		isGregorian = false
	}

	if isGregorian {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			isLeap = true
		}
	} else {
		if year%4 == 0 {
			isLeap = true
		}
	}

	if year == 1918 {
		fmt.Println("26.09.1918")
	} else if isLeap {
		fmt.Printf("12.09.%d\n", year)
	} else {
		fmt.Printf("13.09.%d\n", year)
	}
}
