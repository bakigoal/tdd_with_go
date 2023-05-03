package main

import (
	"os"
	"tdd_with_go/ch_16_maths"
	"time"
)

func main() {
	now := time.Now()
	ch_16_maths.SVGWriter(os.Stdout, now)
}
