package ch_16_maths

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)
	assert.Equal(t, want, got)
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)
	assert.Equal(t, want, got)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), 3 * math.Pi / 2},
		{simpleTime(0, 0, 7), 7 * math.Pi / 30},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)
			assert.Equal(t, c.angle, got)
		})
	}
}

func testName(t time.Time) string {
	return t.Format("15:01:05")
}

func simpleTime(hours int, minutes int, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}
