package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	io := bufio.NewReader(os.Stdin)

	alpha := make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Fscan(io, &alpha[i])
	}

	var input string
	fmt.Fscan(io, &input)
	max := 0
	for i := 0; i < len(input); i++ {
		letter := input[i] - 'a'
		if alpha[letter]> max {
			max = alpha[letter]
		}
	}
	fmt.Printf("%d\n", max * len(input))
}