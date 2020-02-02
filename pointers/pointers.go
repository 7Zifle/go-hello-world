package pointers

import (
	"errors"
	"fmt"
)

var InsufficentError = errors.New("Overdrawn Error")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("BTC: $%d", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Balance() < amount {
		return InsufficentError
	}
	w.balance -= amount
	return nil
}
