package ch_09_mocking

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

func Countdown(out io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
