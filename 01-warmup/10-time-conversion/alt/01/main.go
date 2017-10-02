package main

import (
	"fmt"
	"time"
)

func main() {
	var t string
	fmt.Scanln(&t)

	parsed, _ := time.Parse("3:04:05PM", t)
	fmt.Println(parsed.Format("15:04:05"))
}
