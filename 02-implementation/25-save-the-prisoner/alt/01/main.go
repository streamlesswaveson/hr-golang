package main
import "fmt"

func main() {
	var t int
	var a, b, c int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b, &c)
		ans := (b+c-1)%a
		if ans == 0 {
			fmt.Println(a)
		} else {
			fmt.Println(ans)
		}
	}
}