package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with a name", func(t *testing.T) {
		got := hello("chris")
		want := "Hello, chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("say 'Hello World' when the name string is empty", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
