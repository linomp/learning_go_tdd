package pointers

// lowercase: private outside this package
type Wallet struct {
	balance int
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
func (w *Wallet) Deposit(amount int) {

	// fmt.Printf("\nin method: %v\n", &w.balance)

	// struct pointers get automatically de-referenced
	w.balance += amount
	// manual de-referencing:
	// (*w).balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
	// manual de-referencing:
	// return (*w).balance
}
