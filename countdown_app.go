package main

import (
	"os"
	"tdd_with_go/ch_09_mocking"
	"time"
)

func main() {
	sleeper := &ch_09_mocking.ConfigurableSleeper{ConfigDuration: 2 * time.Second, ConfigSleep: time.Sleep}
	ch_09_mocking.Countdown(os.Stdout, sleeper)
}
