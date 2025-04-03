package ptrserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// In go if a symbole is lowercase (vars, funcs, etc) it is private
type Bitcoin int

type Stringer interface {
	String() string
}

// outside the package cannot use
type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or meethod the arguments are copied
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = w.balance + amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance = w.balance - amount
	return nil
}

// This interface is defined in the fmt package and lets you define how your
// type is printed when used with the %s forrmat string in prints
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
