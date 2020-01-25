package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectGreeting := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}

	t.Run("Say hello with name in french", func(t *testing.T) {
		got := Hello("french", "Jake")
		want := "Bonjour, Jake!"
		assertCorrectGreeting(t, got, want)
	})
	t.Run("Say hello with no name in spanish", func(t *testing.T) {
		got := Hello("spanish", "")
		want := "Hola, buddy!"
		assertCorrectGreeting(t, got, want)
	})
	t.Run("Say hello in english with name", func(t *testing.T) {
		got := Hello("english", "Jake")
		want := "Hello, Jake!"
		assertCorrectGreeting(t, got, want)
	})
}
