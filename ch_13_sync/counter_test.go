package ch_13_sync

import (
	"github.com/stretchr/testify/assert"
	"sync"
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
	t.Run("it run safely concurrently", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup // A WaitGroup waits for a collection of goroutines to finish.
		wg.Add(wantedCount)   // The main goroutine calls 'Add' to set the number of goroutines to wait for.

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // Then each of the goroutines calls 'Done' when finished.
			}()
		}

		wg.Wait() // At the same time, 'Wait' can be used to block until all goroutines have finished.

		assert.Equal(t, wantedCount, counter.Value())
	})
}
