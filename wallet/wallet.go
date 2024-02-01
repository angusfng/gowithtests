package wallet

import (
	"errors"
	"fmt"
)

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

// This var is global to the package
var ErrInsufficientFunds = errors.New("insufficient funds")

// Returns nil on success, otherwise error
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
