package pointers

import (
	"errors"
	"fmt"
)

// making types out of existing ones
// allows to declare methods on types!
type Bitcoin int

// custom error
// var: define values global to the package
var ErrInsufficientFunds = errors.New("Insufficient Funds. Aborted.")

// there is an Stringer interface in the fmt package
// it has the method String() string (gets called on %s)
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// lowercase: private outside this package
type Wallet struct {
	balance Bitcoin
}

// arguments get passed by value! (copies)
// this implementation just modifies  a copy and has no effect
// on the receiver
/*
func (w Wallet) Deposit(amount int) {

	fmt.Printf("\nin method: %v\n", &w.balance)
	w.balance += amount
}
*/

// to modify something on the receiver,
// method must receive a pointer
func (w *Wallet) Deposit(amount Bitcoin) {

	// fmt.Printf("\nin method: %v\n", &w.balance)

	// struct pointers get automatically de-referenced
	w.balance += amount
	// manual de-referencing:
	// (*w).balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
	// manual de-referencing:
	// return (*w).balance
}

// it is idiomatic to return an err
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
