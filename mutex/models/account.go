package models

import (
	"fmt"
	"sync"
)

type Account struct {
	mutex   sync.Mutex
	balance int
}

func NewAccount(init int) *Account {
	return &Account{balance: init}
}

func (a *Account) Deposite(value int, done chan bool) {
	a.mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, a.balance)
	a.balance += value
	a.mutex.Unlock()
	done <- true
}

func (a *Account) Withdraw(value int, done chan bool) {
	a.mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, a.balance)
	a.balance -= value
	a.mutex.Unlock()
	done <- true
}

func (a *Account) Balance() int {
	return a.balance
}
