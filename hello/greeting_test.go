package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to John", func(t *testing.T) {
		got := Hello("John", "en")
		want := "Hello, John!"
		assertEquals(t, got, want)
	})

	t.Run("Скажи привет Джону", func(t *testing.T) {
		got := Hello("Джон", "ru")
		want := "Привет, Джон!"
		assertEquals(t, got, want)
	})

	t.Run("Скажи привет Миру, если не задано имя", func(t *testing.T) {
		got := Hello("", "ru")
		want := "Привет, Мир!"
		assertEquals(t, got, want)
	})

	t.Run("Say hello to World if name is empty", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World!"
		assertEquals(t, got, want)
	})
}

func assertEquals(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
