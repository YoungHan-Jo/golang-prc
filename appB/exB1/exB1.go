package main

import (
	"exB1/bankaccount/bankaccount"
	"fmt"
)

func main() {

	account := bankaccount.NewAccount(10000)
	account.Deposit(1000)

	fmt.Println(account.Balance())

}
