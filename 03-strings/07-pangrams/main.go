package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var in string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in = scanner.Text()
	in = strings.ToLower(in)
	in = strings.Replace(in, " ", "", -1)

	alpha := map[rune]bool{}
	for _, v := range in {
		alpha[v] = true
	}
	if len(alpha) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
