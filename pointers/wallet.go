package pointers

type Wallet struct {
	balance int
}

// a pointer to wallet
func (w *Wallet) Deposit(amount int) {
	// struct pointers automatically deReferenced to
	// (*w).balance += amount
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
