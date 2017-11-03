package main
import "fmt"

func main() {
	var S string
	fmt.Scanf("%s",&S)
	totalWords := 1
	for i:=0; i<len(S); i++ {
		if S[i]>='A' && S[i]<='Z' {
			totalWords++;
		}
	}
	fmt.Println(totalWords)
}