package wallet

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	// Automatically dereferences
	w.balance += amount
} 

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Implements stringer on Bitcoin type so that %s will work with Bitcoin types
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}