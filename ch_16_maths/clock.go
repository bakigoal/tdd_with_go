package ch_16_maths

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondsHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

func secondsHandPoint(t time.Time) Point {
	angle := timeUnitInRadians(t.Second())
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}

func timeUnitInRadians(unit int) float64 {
	return math.Pi * (float64(unit) / 30)
}
