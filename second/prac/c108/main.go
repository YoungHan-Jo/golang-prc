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

	var spotTime []int

	for i := 0; i < n; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		spotTime = append(spotTime, t)
	}

	// var moveTime [][]int
	moveTime := make([][]int, n)
	for i := 0; i < n; i++ {
		moveTime[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")

		for j := 0; j < n; j++ {
			t, _ := strconv.Atoi(strs[j])
			moveTime[i][j] = t
		}
	}

	sc.Scan()

	m, _ := strconv.Atoi(sc.Text())

	var orders []int

	for i := 0; i < m; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		orders = append(orders, t)
	}

	sum := 0

	for i := 0; i < m; i++ {
		sum += spotTime[orders[i]-1]
		if i == m-1 {
			break
		}
		sum += moveTime[orders[i]-1][orders[i+1]-1]
	}

	fmt.Println(sum)
}
