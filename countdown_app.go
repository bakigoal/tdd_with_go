package main

import (
	"os"
	"tdd_with_go/ch_09_mocking"
)

func main() {
	sleeper := &ch_09_mocking.DefaultSleeper{}
	ch_09_mocking.Countdown(os.Stdout, sleeper)
}
