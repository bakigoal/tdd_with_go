package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// a pointer to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// struct pointers automatically deReferenced to
	// (*w).balance += amount
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
