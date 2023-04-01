package ch_09_mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
