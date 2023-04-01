package ch_09_mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := SpySleeper{}

	Countdown(buffer, &sleeper)
	want := `3
2
1
Go!`
	assert.Equal(t, want, buffer.String())
	assert.Equal(t, 3, sleeper.Calls)
}
