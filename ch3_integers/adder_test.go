package main

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {

	t.Run("add 2 integers", func(t *testing.T) {

		got := Sum(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
