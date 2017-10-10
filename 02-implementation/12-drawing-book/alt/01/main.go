package main
import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var n, p int
	if in.Scan() {
		n, _ = strconv.Atoi(in.Text())
	}
	if in.Scan() {
		p, _ = strconv.Atoi(in.Text())
	}

	numpagesfromfront := p / 2
	var numpagesfromback int
	if n % 2 == 0 {
		numpagesfromback = (n - p + 1) / 2
	} else {
		numpagesfromback = (n - p) / 2
	}

	if numpagesfromfront < numpagesfromback {
		fmt.Println(numpagesfromfront)
	} else {
		fmt.Println(numpagesfromback)
	}
}