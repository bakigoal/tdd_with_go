package hello

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("John", "ru")
	expected := "Привет, John"

	if actual != expected {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
