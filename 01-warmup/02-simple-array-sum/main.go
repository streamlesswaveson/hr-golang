package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	_, err := strconv.Atoi(scanner.Text())
	if nil != err {
		panic(err)
	}
	//fmt.Printf("size = %d\n", size)

	scanner.Scan()
	mystrings := strings.Split(scanner.Text(), " ")

	sum := 0
	for _, s := range mystrings {
		myval, err := strconv.Atoi(s)
		if nil != err {
			panic(err)
		}
		sum += myval

	}
	fmt.Printf("%d\n", sum)

}
