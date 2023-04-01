package ch_09_mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	assert.Equal(t, "3", buffer.String())
}
