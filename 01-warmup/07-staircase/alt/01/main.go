package main

import "fmt"
import "strconv"

func main() {
	var h int

	fmt.Scan(&h)

	if h < 1 || h > 100 {
		return
	}

	str := ""
	for i := h; i > 0; i-- {
		str += "#"
		fmt.Printf("%"+strconv.Itoa(h)+"s\n", str)
	}
}