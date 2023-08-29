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

	n, _ := strconv.Atoi(sc.Text())

	min := 100
	max := 0

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")

		start, _ := strconv.Atoi(strs[0])
		move, _ := strconv.Atoi(strs[1])
		end, _ := strconv.Atoi(strs[2])

		sum := start + move + 24 - end

		if min > sum {
			min = sum
		}

		if max < sum {
			max = sum
		}
	}

	fmt.Println(min)
	fmt.Println(max)
}
