package main
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	bio := bufio.NewReader(os.Stdin)
	line, _ := bio.ReadString('\n')
	count, _ := strconv.Atoi(strings.TrimSpace(line))
	for i := 0; i < count; i++ {
		line, _ := bio.ReadString('\n')
		fields := strings.Fields(line)
		low, _ := strconv.Atoi(fields[0])
		high, _ := strconv.Atoi(fields[1])
		min := int(math.Sqrt(float64(low)))
		max := int(math.Sqrt(float64(high)))
		total := max - min
		if int(math.Pow(float64(min), 2)) == low {
			total++
		}
		fmt.Println(total)
	}
}