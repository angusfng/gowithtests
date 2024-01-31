package wallet

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

// Implements stringer on Bitcoin type so that %s will work with Bitcoin types
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Pointer to wallet because we want to modify it
func (w *Wallet) Deposit(amount Bitcoin) {
	// Automatically dereferences
	w.balance += amount
}

// Passes in a copy
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
