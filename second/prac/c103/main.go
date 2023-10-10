package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	interval int
	action   string
}

func (a *Action) activeAction(sec, sum int) bool {
	if sec%a.interval == 0 {
		if sum > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%s", a.action)
		return true
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	n, _ := strconv.Atoi(strs[0])
	m, _ := strconv.Atoi(strs[1])

	var actions []Action

	for i := 0; i < m; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")
		interval, _ := strconv.Atoi(strs[0])

		actions = append(actions, Action{interval, strs[1]})
	}

	for i := 1; i <= n; i++ {
		sum := 0
		for _, a := range actions {
			if a.activeAction(i, sum) {
				sum += 1
			}
		}
		if sum == 0 {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}

}
