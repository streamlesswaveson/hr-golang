package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var raw string
	fmt.Scan(&raw)

	parts := strings.Split(raw, ":")

	if strings.Contains(parts[2], "AM") {
		if "12" == parts[0] {
			parts[0] = "00"
		}
		parts[2] = strings.Replace(parts[2], "AM", "", 1)
	} else {
		hour, _ := strconv.Atoi(parts[0])
		if hour < 12 {
			hour += 12
			parts[0] = fmt.Sprintf("%d", hour)
		}
		parts[2] = strings.Replace(parts[2], "PM", "", 1)
	}

	returned := strings.Join(parts, ":")

	fmt.Println(returned)
}
