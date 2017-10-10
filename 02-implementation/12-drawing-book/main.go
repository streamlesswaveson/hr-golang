package main

import "fmt"
//https://www.hackerrank.com/challenges/drawing-book/editorial
func main() {
	var totalPages, targetPage int
	fmt.Scan(&totalPages)
	fmt.Scan(&targetPage)

	var left2Right int
	if targetPage%2 == 0 {
		left2Right = (targetPage + 2) / 2
	} else {
		left2Right = (targetPage + 1) / 2
	}
	left2Right = left2Right - 1

	var right2Left int
	if (totalPages%2 == 0 && targetPage%2 == 0) || (totalPages%2 != 0 && targetPage%2 != 0) {
		right2Left = (totalPages - targetPage) / 2
	} else if totalPages%2==0 {
		right2Left = (totalPages - (targetPage - 1)) / 2
	} else if targetPage%2==0 {
		right2Left = (totalPages - (targetPage+1) ) / 2
	}
	//fmt.Println(left2Right, right2Left)
	if left2Right < right2Left {
		fmt.Println(left2Right)
	} else {
		fmt.Println(right2Left)
	}
}
