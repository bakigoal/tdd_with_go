package ch_13_sync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("inc the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})
}
