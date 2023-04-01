package main

import (
	"os"
	"tdd_with_go/ch_09_mocking"
	"time"
)

func main() {
	sleeper := &ch_09_mocking.ConfigurableSleeper{2 * time.Second, time.Sleep}
	ch_09_mocking.Countdown(os.Stdout, sleeper)
}
