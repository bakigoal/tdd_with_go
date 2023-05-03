package ch_16_maths

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}

func timeToRadians(unit int) float64 {
	return math.Pi * (float64(unit) / 30)
}

func SVGWriter(w io.Writer, tm time.Time) {
	io.WriteString(w, svgStart)
	newLine(w)

	io.WriteString(w, bezel)
	newLine(w)

	HourHand(w, tm)
	newLine(w)

	MinuteHand(w, tm)
	newLine(w)

	SecondHand(w, tm)
	newLine(w)

	io.WriteString(w, svgEnd)
}

func newLine(w io.Writer) {
	io.WriteString(w, "\n")
}

func HourHand(w io.Writer, t time.Time) {
	p := getPoint(t.Hour()*5, hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}

func MinuteHand(w io.Writer, t time.Time) {
	p := getPoint(t.Minute(), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}

func SecondHand(w io.Writer, t time.Time) {
	p := getPoint(t.Second(), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func getPoint(time int, length float64) Point {
	p := angleToPoint(timeToRadians(time))
	p = Point{p.X * length, p.Y * length}             // scale
	p = Point{p.X, -p.Y}                              // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} // translate
	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
	 height="100%"
     viewBox="0 0 300 300">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
