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
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var result [4]int
	result[3] = math.MaxInt64

	for i := 0; i < n; i++ {
		sc.Scan()
		str := strings.Split(sc.Text(), " ")
		if i == 0 {
			result[0], _ = strconv.Atoi(str[0])
		}
		if i == n-1 {
			result[1], _ = strconv.Atoi(str[1])
		}

		if v, _ := strconv.Atoi(str[2]); v > result[2] {
			result[2] = v
		}

		if v, _ := strconv.Atoi(str[3]); v < result[3] {
			result[3] = v
		}
	}

	fmt.Printf("%d %d %d %d", result[0], result[1], result[2], result[3])
}
