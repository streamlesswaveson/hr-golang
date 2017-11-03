package main
import "fmt"

func check(s string) bool {
	n:=len(s)
	if n==1 {
		return false
	}
	m:=map[byte]bool{}
	for i:=0; i<n-1; i++ {
		if s[i]==s[i+1] {
			return false
		}
		m[s[i]]=true
		if len(m) > 2 {
			return false
		}
	}
	return true
}

func removeAllExcept(s string, a, b byte) string {
	res := ""
	for i := range s {
		if s[i]==a || s[i]== b {
			res += string(s[i])
		}
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	s:=""
	fmt.Scanln(&s)
	if check(s) {
		fmt.Println(n)
		return
	}
	m:=map[byte]bool{}
	for i:=range s {
		m[s[i]] = true
	}
	max := 0
	for a:=range m {
		for b := range m {
			if a==b {
				continue
			}
			t:=removeAllExcept(s, a, b)
			if check(t) {
				if max < len(t) {
					max = len(t)
				}
			}
		}
	}
	fmt.Println(max)
}