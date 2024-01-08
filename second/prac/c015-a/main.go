package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Money struct {
	amount int
}

type Recipe struct {
	day   string
	price Money
}

func (r *Recipe) getPoint() *Point {
	if strings.Contains(r.day, "3") {
		return &Point{r.price.amount * 3 / 100}
	}

	if strings.Contains(r.day, "5") {
		return &Point{r.price.amount * 5 / 100}
	}

	return &Point{r.price.amount * 1 / 100}
}

type Point struct {
	amount int
}

func (p *Point) plusPoints(newPoint *Point) Point {
	return Point{p.amount + newPoint.amount}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	totalPoints := Point{0}

	for i := 0; i < n; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")
		day := strs[0]
		price, _ := strconv.Atoi(strs[1])

		recipe := Recipe{day, Money{price}}

		totalPoints = totalPoints.plusPoints(recipe.getPoint())
	}

	fmt.Println(totalPoints.amount)
}
