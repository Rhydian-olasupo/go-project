package iteration

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	t.Run("test for no value supplied", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "Input a number above 0"
		assertCorrectTest(t, got, want)

	})

	t.Run("test when a value is supplied", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		assertCorrectTest(t, got, want)
	})
}

func assertCorrectTest(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	word := Repeat("a", 10)
	fmt.Println(word)
	// Output: aaaaaaaaaa
}
