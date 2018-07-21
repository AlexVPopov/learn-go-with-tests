package main

func main() {}

// Bitcoin is a bitcoin
type Bitcoin int

// Wallet holds bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit adds bitcoins to the balance of a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the number of bitcoins in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
