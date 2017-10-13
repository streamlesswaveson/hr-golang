package main
import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d", &n)
	score := make([]int, n)
	scoreLen := 0
	var nextScore int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nextScore)
		if scoreLen == 0 || nextScore != score[scoreLen - 1] {
			score[scoreLen] = nextScore
			scoreLen++
		}
	}
	fmt.Scanf("%d", &m)
	curPos := scoreLen - 1
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &nextScore)
		for curPos >= 0 && nextScore >= score[curPos] {
			curPos--
		}
		fmt.Println(curPos + 2)
	}
}