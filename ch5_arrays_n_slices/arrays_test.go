package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t *testing.T, got []int, want []int) {
		t.Helper()

		// beware, not type-safe!
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("sum contents of a slice", func(t *testing.T) {

		// fixed-size arrays, two ways:
		// [N]type{value1, value2, ..., valueN}
		// [...]type{value1, value2, ..., valueN}
		// size is encoded in type!
		// [4]int diff type from [3]int
		nums := []int{1, 2, 3, 4, 5, 6} // slice
		got := SumCollection(nums)
		want := 21
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, nums)
		}

	})

	t.Run("sum contents of all slices passed", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("sum tails of all slices passed", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3}, []int{2, 9, 1})
		want := []int{5, 10}

		checkSums(t, got, want)
	})

	t.Run("safely sum edge-case slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{7}, []int{3, 4, 5})
		want := []int{0, 7, 9}

		checkSums(t, got, want)
	})
}
