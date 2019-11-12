package main

import "testing"

// always refactor tests when possible
// tests should be clear specifications of what the code needs to do.
func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got string, want string) {
		t.Helper() // so when it fails it doesn't report this line but where this fn iscalled
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// a subtest
	t.Run("say hello when receiving a name", func(t *testing.T) {
		got := Hello("Vio")
		want := "Hello, Vio"
		assertMessage(t, got, want)
	})

	t.Run("defaults to hello world upon an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertMessage(t, got, want)
	})
}

/*
Run with: go test

Rules:
	xxx_test.go
	start with Test...
	arg: t *testing.T  -> "hook" into the testing framework

fmt placeholders:
	https://golang.org/pkg/fmt/#hdr-Printing
	%q -> double quotes

godoc in Go 1.13+
	go get golang.org/x/tools/cmd/godoc


*/
