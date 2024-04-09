package main

type Account struct {
	balance int
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func (c *Account) simulate(value int) int {
	c.balance += value
	println(c.balance)
	return c.balance
}

func main() {
	account := Account{balance: 100}
	account.simulate(200)
	println(account.balance)
}
