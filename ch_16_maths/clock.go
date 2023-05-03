package ch_16_maths

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{}
}

func SecondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}
