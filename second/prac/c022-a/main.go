package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Amount struct {
	amount float64
}

func (a *Amount) calPercent(percent int) Amount {
	return Amount{a.amount / 100 * float64(percent)}
}

type Product struct {
	remains Amount
}

func (p *Product) sell(soldPer int) {
	p.remains = p.remains.calPercent(100 - soldPer)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	origin, _ := strconv.Atoi(strs[0])
	firstPer, _ := strconv.Atoi(strs[1])
	secondPer, _ := strconv.Atoi(strs[2])

	product := Product{Amount{float64(origin)}}

	product.sell(firstPer)
	product.sell(secondPer)

	fmt.Println(product.remains.amount)

}
