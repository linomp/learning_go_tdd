package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	// controled access via methods

	// fmt.Printf("\nin test: %v\n", &wallet.balance)
	got := wallet.Balance()

	want := 10

	if got != want {
		t.Errorf("got %d want %d!", got, want)
	}

}
