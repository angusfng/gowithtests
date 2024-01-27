package helloworld

import "testing"

func TestHello(t *testing.T) {
	// Subtests
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Angus", "")
		want := "Hello, Angus"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Angus", "Spanish")
		want := "Hola, Angus"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Angus", "French")
		want := "Bonjour, Angus"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}