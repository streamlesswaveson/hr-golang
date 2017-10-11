package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	q := readInt()
	for t := 0; t < q; t++ {
		x, y, z := readInt(), readInt(), readInt()
		da, db := abs(x-z), abs(y-z)
		res := "Mouse C"
		if da < db {
			res = "Cat A"
		} else if da > db {
			res = "Cat B"
		}
		fmt.Println(res)
	}

}

// Helpers

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

// Math

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}

// Sort

type Ints64 []int64

func (a Ints64) Len() int           { return len(a) }
func (a Ints64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints64) Less(i, j int) bool { return a[i] < a[j] }

/* Sort tempalte
func (a ) Len() int           { return len(a) }
func (a ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ) Less(i, j int) bool { return  }
*/