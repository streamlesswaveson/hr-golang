package main

import "fmt"

func main() {
	var (
		N, M, lastScore int
	)
	fmt.Scan(&N)

	all := make([]int, N)
	fmt.Scan(&all[0])
	//leader := map[int]int{}
	currentRank := 1
	//leader[currentRank]++
	rankByScore := map[int]int{}
	rankByScore[currentRank] = all[0]

	lastScore = all[0]
	for i := 1; i < N; i++ {
		fmt.Scan(&all[i])
		if all[i] == lastScore {
			//leader[currentRank]++
		} else {
			currentRank++
			//leader[currentRank]++
			lastScore = all[i]
			rankByScore[currentRank] = lastScore
		}
	}
	//fmt.Println(leader)
	//fmt.Println(rankByScore)

	fmt.Scan(&M)
	alice := make([]int, M)

	for i := 0; i < M; i++ {
		fmt.Scan(&alice[i])
		// last on the leaderboard
		if alice[i] < lastScore {
			p(currentRank + 1, alice[i])
			continue
		}
		if alice[i] == lastScore {
			p(currentRank, alice[i])
			continue
		}

		for {
			currentRank--
			nextScore := rankByScore[currentRank]
			if alice[i] == nextScore {
				p(currentRank, alice[i])
				break
			}
			if alice[i] < nextScore {
				p(currentRank + 1, alice[i])
				currentRank++
				break
			}
			if alice[i] > nextScore && currentRank <= 1 {
				p(1, alice[i])
				break
			}
		}

	}

}

func p(rank int, score int) {
	fmt.Println(rank)
	//fmt.Println(rank, score)
}
