package main
import "fmt"

func main() {
	var a = make([]int, 3)
	var b = make([]int, 3)

	// fmt.Scan takes varargs
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])

	as := 0
	bs := 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			as++
		} else if a[i] < b[i] {
			bs++
		}
	}

	fmt.Println(as, bs)
}