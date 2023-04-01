package ch_01_hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say ch_01_hello to John", func(t *testing.T) {
		got := Hello("John", "en")
		want := "Hello, John!"
		assertEquals(t, got, want)
	})
	t.Run("Скажи привет Джон", func(t *testing.T) {
		got := Hello("Джон", "ru")
		want := "Привет, Джон!"
		assertEquals(t, got, want)
	})
	t.Run("Скажи привет Мир", func(t *testing.T) {
		got := Hello("", "ru")
		want := "Привет, Мир!"
		assertEquals(t, got, want)
	})
	t.Run("Say ch_01_hello World", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World!"
		assertEquals(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola, Elodie!"
		assertEquals(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "fr")
		want := "Bonjour, Elodie!"
		assertEquals(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB
// which is an interface that *testing.T and *testing.B
func assertEquals(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call
	// rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
