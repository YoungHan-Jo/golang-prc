package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Recipe struct {
	day   string
	price int
}

func (r *Recipe) getPoint() int {
	var percent = 0.01
	switch r.day {
	case "3", "13", "23", "30", "31":
		percent = 0.03
	case "5", "15", "25":
		percent = 0.05
	}

	return int(math.Floor(float64(r.price) * percent))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	sum := 0

	for i := 0; i < n; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")

		price, _ := strconv.Atoi(strs[1])

		recipe := Recipe{strs[0], price}

		sum += recipe.getPoint()
	}

	fmt.Println(sum)

}
