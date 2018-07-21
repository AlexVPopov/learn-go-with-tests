package main

import (
	"errors"
	"fmt"
)

func main() {}

// Stringer stringifies the implementor
type Stringer interface {
	String() string
}

// Bitcoin is a bitcoin
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet holds bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit adds bitcoins to the balance of a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes bitcoins from the balance of wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount

	return nil
}

// Balance returns the number of bitcoins in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
