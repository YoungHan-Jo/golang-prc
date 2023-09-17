package bankaccount

type Account interface {
	Withdraw(money int) int
	Deposit(money int)
	Balance() int
}

func NewAccount(balance int) Account {
	return &innerAccount{balance}
}

type innerAccount struct {
	balance int
}

func (a *innerAccount) Withdraw(money int) int {
	a.balance -= money
	return a.balance
}

func (a *innerAccount) Deposit(money int) {
	a.balance += money
}

func (a *innerAccount) Balance() int {
	return a.balance
}
