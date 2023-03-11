package hello

import "testing"

func TestHello(t *testing.T) {
	assertEquals(t, Hello("John", "ru"), "Привет, John!")
	assertEquals(t, Hello("John", "en"), "Hello, John!")
	assertEquals(t, Hello("", "en"), "Hello, world!")
	assertEquals(t, Hello("", "ru"), "Привет, мир!")
}

func assertEquals(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
