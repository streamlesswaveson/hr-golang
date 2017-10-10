package main
import "fmt"

func main() {
	birdCount := make(map[int]int, 5)
	var n, maxBird, bird int
	maxBird = 1
	fmt.Scanf("%d", &n)
	for i := 0;i < n;i++ {
		fmt.Scanf("%d", &bird)
		birdCount[bird] += 1
		if(birdCount[bird] > birdCount[maxBird]) {
			maxBird = bird
		} else if(birdCount[bird] == birdCount[maxBird] && bird < maxBird) {
			maxBird = bird
		}
	}
	fmt.Println(maxBird)
}