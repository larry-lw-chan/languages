package wallet

import "errors"

type Banking interface {
	Deposit()
	Balance()
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	w.balance += bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin < w.balance {
		w.balance -= bitcoin
		return nil
	} else {
		return ErrInsufficientFunds
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
