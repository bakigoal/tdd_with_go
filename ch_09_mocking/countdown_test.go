package ch_09_mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	want := `3
2
1
Go!`
	assert.Equal(t, want, buffer.String())
}
