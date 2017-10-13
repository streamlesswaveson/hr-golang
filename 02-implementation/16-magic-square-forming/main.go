package main

import (
	"fmt"
)

// you're on the wrong track - revisit later
func main() {
	square := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&square[i][j])
		}
	}
	fmt.Println(square)
	fmt.Println(isMagic(square))

}

func isMagic(s [3][3]int ) bool {
	seen := map[int]bool{}

	expectedSum := -1

	for i:=0; i<3; i++ {
		var sum int
		for j:=0; j<3; j++ {
			v := s[i][j]
			if seen[v] {
				return false
			}
			sum += v
			seen[v] = true
		}
		if expectedSum < 0 {
			expectedSum = sum
		} else {
			if sum != expectedSum {
				return false
			}
		}
	}

	for i:=0; i<3; i++{
		if s[0][i] + s[1][i] + s[2][i] != expectedSum {
			return false
		}
	}

	if s[0][0] + s[1][1] + s[2][2] != expectedSum {
		return false
	}

	if s[0][2] + s[1][1] + s[2][0] != expectedSum {
		return false
	}
	return true
}
