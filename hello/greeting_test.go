package hello

import "testing"

func TestHelloRu(t *testing.T) {
	assertHello(t, "John", "ru", "Привет, John!")
}

func TestHelloEn(t *testing.T) {
	assertHello(t, "John", "en", "Hello, John!")
}

func assertHello(t *testing.T, name, lang, want string) {
	got := Hello(name, lang)

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
