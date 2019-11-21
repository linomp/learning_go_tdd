package iteration

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 3))
	//Output: bbb
}
