package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cooking(amount float64, percent int) float64 {
	return amount / 100 * float64(100-percent)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	m, _ := strconv.Atoi(strs[0])
	p, _ := strconv.Atoi(strs[1])
	q, _ := strconv.Atoi(strs[2])

	fmt.Println(cooking(cooking(float64(m), p), q))
}
