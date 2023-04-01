package ch_09_mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "SleepFn"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})
		want := `3
2
1
Go!`
		assert.Equal(t, want, buffer.String())
	})
	t.Run("SleepFn before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{write, sleep, write, sleep, write, sleep, write}

		assert.Equal(t, want, spySleepPrinter.Calls)
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	assert.Equal(t, sleepTime, spyTime.durationSlept)
}
