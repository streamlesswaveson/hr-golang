package main
import "fmt"

func main() {
	var n, tmp, sum int;
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}
	fmt.Println(sum)
}