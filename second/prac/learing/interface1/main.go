package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Customer interface {
	OrderFood(*int)
	OrderAlcohol(*int)
	OrderSoftdrink(*int)
	GetAmmount() *int
}

func NewAdult() *adult {
	return &adult{false, 0}
}

func NewKid() *kid {
	return &kid{0}
}

type adult struct {
	isDrink bool
	ammount int
}

func (a *adult) OrderFood(price *int) {
	a.ammount += *price
	if a.isDrink {
		a.ammount -= 200
	}
}

func (a *adult) OrderAlcohol(price *int) {
	a.ammount += *price
}

func (a *adult) OrderSoftdrink(price *int) {
	a.ammount += *price
}

func (a *adult) GetAmmount() *int {
	return &a.ammount
}

type kid struct {
	ammount int
}

func (k *kid) OrderFood(price *int) {
	k.ammount += *price
}

func (k *kid) OrderAlcohol(price *int) {

}

func (k *kid) OrderSoftdrink(price *int) {
	k.ammount += *price
}

func (k *kid) GetAmmount() *int {
	return &k.ammount
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(strs[0])
	m, _ := strconv.Atoi(strs[1])

	var customers []Customer

	for i := 0; i < n; i++ {
		sc.Scan()
		old, _ := strconv.Atoi(sc.Text())
		if old >= 20 {
			customers = append(customers, NewAdult())
			continue
		}
		customers = append(customers, NewKid())
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		index, _ := strconv.Atoi(strs[0])
		price, _ := strconv.Atoi(strs[2])

		switch strs[1] {
		case "food":
			customers[index-1].OrderFood(&price)
		case "alcohol":
			customers[index-1].OrderAlcohol(&price)
		case "softdrink":
			customers[index-1].OrderSoftdrink(&price)
		}
	}

	for _, c := range customers {
		fmt.Println(*c.GetAmmount())
	}

}
