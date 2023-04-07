package ch_11_select

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeoutDuration = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultTimeoutDuration)
}

func ConfigurableRacer(a, b string, duration time.Duration) (winner string, err error) {
	// select lets you wait on multiple channels
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		// “Sometimes you'll want to include time.After to prevent blocking forever”
	case <-time.After(duration):
		return "", fmt.Errorf("timeout waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
