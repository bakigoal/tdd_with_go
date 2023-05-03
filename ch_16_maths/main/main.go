package main

import (
	"fmt"
	"io"
	"os"
	"tdd_with_go/ch_16_maths"
	"time"
)

func main() {
	now := time.Now()
	secondHand := ch_16_maths.SecondHand(now)

	io.WriteString(os.Stdout, svgStart)
	io.WriteString(os.Stdout, bezel)
	io.WriteString(os.Stdout, secondHandTag(secondHand))
	io.WriteString(os.Stdout, svgEnd)
}

func secondHandTag(p ch_16_maths.Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" 
	style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
	 height="100%"
     viewBox="0 0 300 300">
`

const bezel = `
<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>
`

const svgEnd = `
</svg>`
