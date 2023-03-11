package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("John", "ru")
	want := "Привет, John!"

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
