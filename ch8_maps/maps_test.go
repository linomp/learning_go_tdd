package main

import "testing"

func TestSearch(t *testing.T) {

	// key type has to be a comparable type
	// map[key type]value type
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"
	assertStrings(t, got, want)
}

func TestSearchCustom(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrUnknownWord)

	})

}

func assertError(t *testing.T, got error, want error) {

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
}

func TestUpdate(t *testing.T) {
}

func TestDelete(t *testing.T) {
}
