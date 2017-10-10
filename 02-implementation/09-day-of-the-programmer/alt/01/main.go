package main

import (
	"fmt"
	"time"
)

func russiaProgramerDay(year int) string {
	if year < 1918 {
		if year%4 == 0 {
			return fmt.Sprintf("12.09.%d", year)
		} else {
			return fmt.Sprintf("13.09.%d", year)
		}
	}

	t := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	days := 255 * time.Hour * 24
	if year == 1918 {
		days = (256 - 32) * time.Hour * 24
		t = time.Date(year, time.February, 14, 0, 0, 0, 0, time.UTC)
	}

	return t.Add(days).Format("02.01.2006")
}

func main() {
	var d int
	fmt.Scan(&d)
	fmt.Println(russiaProgramerDay(d))
}
