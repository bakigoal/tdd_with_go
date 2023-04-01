package ch_08_dependency_injection

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John Snow")

	got := buffer.String()
	want := "Hello, John Snow!"
	assert.Equal(t, want, got)
}
