package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs[0])
	down, _ := strconv.Atoi(strs[1])
	up, _ := strconv.Atoi(strs[2])

	count := 0
	sum := 0

	for i := 1; i <= n; i++ {
		sc.Scan()
		stock, _ := strconv.Atoi(sc.Text())

		if i == n && count > 0 {
			sum += stock * count
			break
		}

		if stock <= down {
			count++
			sum -= stock
		}
		if stock >= up {
			sum += stock * count
			count = 0
		}
	}

	fmt.Println(sum)
}
