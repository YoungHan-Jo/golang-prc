package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			for {
				account.DepositAndWithdraw()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func (account *Account) DepositAndWithdraw() {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
