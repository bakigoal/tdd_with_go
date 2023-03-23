package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// a pointer to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// struct pointers automatically deReferenced to
	// (*w).balance += amount
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("not enough money")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
