package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	first  int
	second int
}

func (p *Card) battle(c *Card) string {
	if p.first > c.first {
		return "High"
	}
	if p.first < c.first {
		return "Low"
	}

	if p.second < c.second {
		return "High"
	}

	if p.second > c.second {
		return "Low"
	}
	return "Low"
}

func makeCard(sc *bufio.Scanner) *Card {
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	x, _ := strconv.Atoi(strs[0])
	y, _ := strconv.Atoi(strs[1])

	return &Card{x, y}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	parentCard := makeCard(sc)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var results []string

	for i := 0; i < n; i++ {
		child := makeCard(sc)
		results = append(results, parentCard.battle(child))
	}

	for _, v := range results {
		fmt.Println(v)
	}

}
