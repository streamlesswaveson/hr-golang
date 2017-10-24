package main

import "fmt"

func main() {

	fmt.Println(solve())
}

func solve() int {
	var (
		actualDay, actualMonth, actualYear       int
		expectedDay, expectedMonth, expectedYear int
	)
	fmt.Scan(&actualDay, &actualMonth, &actualYear)
	fmt.Scan(&expectedDay, &expectedMonth, &expectedYear)

	if actualYear > expectedYear {
		return 10000
	}

	if actualYear < expectedYear {
		return 0
	}

	if actualMonth < expectedMonth {
		return 0
	}
	if actualMonth > expectedMonth {
		return (actualMonth - expectedMonth) * 500
	}
	if actualDay > expectedDay {
		return (actualDay - expectedDay) * 15
	}

	return 0

}
