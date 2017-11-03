package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		N int
		S string
	)
	fmt.Scan(&N, &S)

	if isValid(S) {
		fmt.Println(len(S))
		return
	}

	distinct := map[byte]bool{}

	for i := 0; i < len(S); i++ {
		distinct[S[i]] = true
	}
	debug(distinct)
	distinctPairs := [][]byte{}
	keys := []byte{}
	for k := range distinct {
		keys = append(keys, k)
	}
	debug(keys)
	for i := 0; i < len(keys); i++ {
		curr := keys[i]
		for j := i + 1; j < len(keys); j++ {
			arr := []byte{curr, keys[j]}
			debug("adding", arr)
			distinctPairs = append(distinctPairs, arr)
		}
	}
	debug(distinctPairs)

	maximumLength := 0
	for _, pair := range distinctPairs {
		testString := S

		for k := range distinct {
			if k != pair[0] && k != pair[1] {
				testString = strings.Replace(testString, string(k), "", -1)
			}
		}
		debug("testString = ", testString)
		if isValid(testString) && len(testString) > maximumLength {
			maximumLength = len(testString)
		}
	}
	fmt.Println(maximumLength)
}

func debug(i ...interface{}) {
	//fmt.Println(i)
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	first, second := s[0], s[1]
	if first == second {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == first && s[i+1] == second {
			i++
			continue
		} else {
			return false
		}
	}
	if len(s)%2 == 1 {
		if s[len(s)-1] != first {
			return false
		}
	}
	return true

}
