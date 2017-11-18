package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i:=0; i<N; i++ {
		fmt.Println(solveIt())
	}
}

func solveIt() string {
	var S string
	fmt.Scan(&S)

	if len(S) == 1 {
		return "NO"
	}
	if S[0] == '0' {
		return "NO"
	}

	offset := 1
	for offset < len(S) {
		//fmt.Println("offset=",offset)
		cand := string(S[len(S)-offset:])
		//fmt.Println("cand=",cand)
		if cand[0] == '0' {
			offset++
			continue
		}
		v,e := strconv.Atoi(cand)
		if e != nil {
			offset++
			continue
		}
		remainder := string(S[0:len(S)-len(cand)])
		if len(cand) > len(remainder)+1 {
			return "NO"
		}
		answer, isValid := solve(v, remainder)
		if isValid {
			return answer
		} else {
			offset++
		}
	}
	return "NO"

}

func solve(a int, b string) (string,bool) {
	//fmt.Println("a=",a,"b=",b)
	if a <= 1 {
		return "NO", false
	}
	expected := a - 1
	sExpected := strconv.Itoa(expected)
	if b == sExpected {
		return fmt.Sprintf("YES %s", sExpected),true
	}
	if len(sExpected) > len(b) {
		return "NO", false
	}
	check := b[len(b)-len(sExpected):]
	//fmt.Println("check=",check)
	if check == sExpected {
		return solve(expected, b[0:len(b)-len(sExpected)])
	}
	return "NO",false

}