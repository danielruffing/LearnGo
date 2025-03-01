package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Dan", "Eng")
		want := "Hello, Dan"

		assertCorretMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorretMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorretMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jeeves", "French")
		want := "Bonjour, Jeeves"
		assertCorretMessage(t, got, want)
	})
}

func assertCorretMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
