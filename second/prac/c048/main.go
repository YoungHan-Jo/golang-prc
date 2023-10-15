package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func discount(price int, percent float64) int {
	return int(float64(price) / 100.0 * percent)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")
	price, _ := strconv.Atoi(strs[0])
	percent, _ := strconv.Atoi(strs[1])

	sum := price
	for price != 0 {
		price = discount(price, float64(100-percent))
		sum += price
	}
	fmt.Println(sum)
}
