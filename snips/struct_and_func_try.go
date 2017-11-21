// A good resource for using struct for a "Fund" project:
// https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial
package main

import "fmt"

type Account struct {
	balance int
}

func NewAccount(initialBalance int) *Account {
	return &Account{
		balance: initialBalance,
	}
}

func (account *Account) Withdraw(amount int) {
	account.balance -= amount
}

func main() {
	account := NewAccount(100)
	account.Withdraw(10)

	fmt.Println("Account balance:", account.balance)
}