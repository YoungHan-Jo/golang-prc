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

func (m *Money) plus(money *Money) Money {
	return Money{m.amount + money.amount}
}

func (m *Money) minus(money *Money) Money {
	return Money{m.amount - money.amount}
}

type Customer struct {
	age       int
	isDrunken bool
	bill      Money
}

func (c *Customer) order(item string, price int) {
	if item == "0" {
		if c.isKid() {
			return
		}
		c.plusBill(500)
		c.orderAlcohol()
	}
	if item == "alcohol" {
		if c.isKid() {
			return
		}
		c.plusBill(price)
		c.orderAlcohol()
	}
	if item == "softdrink" {
		c.plusBill(price)
	}
	if item == "food" {
		c.plusBill(price)
		c.discount()
	}
}

func (c *Customer) isKid() bool {
	return c.age < 20
}

func (c *Customer) discount() {
	if c.isDrunken {
		c.bill = c.bill.minus(&Money{200})
	}
}

func (c *Customer) orderAlcohol() {
	c.isDrunken = true
}

func (c *Customer) plusBill(price int) {

	c.bill = c.bill.plus(&Money{price})
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	strs := strings.Split(sc.Text(), " ")

	customCount, _ := strconv.Atoi(strs[0])
	orderCount, _ := strconv.Atoi(strs[1])

	var customers []Customer

	for i := 1; i <= customCount; i++ {
		sc.Scan()
		age, _ := strconv.Atoi(sc.Text())

		customer := Customer{age, false, Money{}}
		customers = append(customers, customer)
	}

	for i := 0; i < orderCount; i++ {
		sc.Scan()

		strs := strings.Split(sc.Text(), " ")

		num, _ := strconv.Atoi(strs[0])
		item := strs[1]

		if len(strs) > 2 {
			price, _ := strconv.Atoi(strs[2])
			customers[num-1].order(item, price)
		} else {
			customers[num-1].order(item, 500)
		}
	}

	for _, customer := range customers {
		fmt.Println(customer.bill.amount)
	}

}
