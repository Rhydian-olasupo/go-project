package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with a name string when language string is empty", func(t *testing.T) {
		got := hello("chris", "")
		want := "Hello, chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello world' when the name and language string is empty", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := hello("Eldie", "Spanish")
		want := "Hola, Eldie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := hello("", "French")
		want := "Bonjour, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
