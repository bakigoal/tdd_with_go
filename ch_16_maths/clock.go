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
	return timeUnitInRadians(t.Second())
}

func SecondsHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}

func timeUnitInRadians(unit int) float64 {
	return math.Pi * (float64(unit) / 30)
}
