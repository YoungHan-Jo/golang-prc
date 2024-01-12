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

func (c *Card) contract(card *Card) string {

	if c.first < card.first {
		return "Low"
	}

	if c.first > card.first {
		return "High"
	}

	if c.second > card.second {
		return "Low"
	}
	return "High"
}

func makeCard(strs *[]string) *Card {
	x, _ := strconv.Atoi((*strs)[0])
	y, _ := strconv.Atoi((*strs)[1])

	return &Card{x, y}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	first, _ := strconv.Atoi(strs[0])
	second, _ := strconv.Atoi(strs[1])

	origin := Card{first, second}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()

		strs = strings.Split(sc.Text(), " ")

		card := makeCard(&strs)

		fmt.Printf("%s\n", origin.contract(card))
	}
}
