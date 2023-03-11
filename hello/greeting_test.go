package hello

import "testing"

func TestHelloRu(t *testing.T) {
	got := Hello("John", "ru")
	want := "Привет, John!"

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}

func TestHelloEn(t *testing.T) {
	got := Hello("John", "en")
	want := "Hello, John!"

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
