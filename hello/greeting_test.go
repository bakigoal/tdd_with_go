package hello

import "testing"

func TestHello(t *testing.T) {
	assertHello(t, "John", "ru", "Привет, John!")
	assertHello(t, "John", "en", "Hello, John!")
}

func assertHello(t *testing.T, name, lang, want string) {
	got := Hello(name, lang)

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
