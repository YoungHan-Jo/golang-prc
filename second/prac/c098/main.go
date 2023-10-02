package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Man struct {
	ballCount int
}

func (m *Man) passBall(amount *int) int {
	passedCount := *amount
	if m.ballCount < *amount {
		passedCount = m.ballCount
		m.ballCount = 0
	} else {
		m.ballCount -= *amount
	}
	return passedCount
}

func (m *Man) getBall(amount *int) {
	m.ballCount += *amount
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var men []Man

	for i := 0; i < n; i++ {
		sc.Scan()
		ballCount, _ := strconv.Atoi(sc.Text())
		men = append(men, Man{ballCount})
	}

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	for i := 0; i < m; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		prev, _ := strconv.Atoi(strs[0])
		next, _ := strconv.Atoi(strs[1])
		amount, _ := strconv.Atoi(strs[2])

		passedBall := men[prev-1].passBall(&amount)
		men[next-1].getBall(&passedBall)
	}

	for _, v := range men {
		fmt.Println(v.ballCount)
	}
}
