package iteration

import "testing"

func TestIter(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	assertCorrectTest(t, got, want)
}

func assertCorrectTest(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
