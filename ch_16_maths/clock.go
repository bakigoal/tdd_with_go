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
	p := SecondsHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scale
	p = Point{p.X, -p.Y}            // flip
	p = Point{p.X + 150, p.Y + 150} // translate
	return p
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
