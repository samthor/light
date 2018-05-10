package light

import (
	"time"
)

// Color is a triple of r,g,b.
type Color [3]byte

// Equal returns whether the two colors are equal.
func (c *Color) Equal(other *Color) bool {
	if c == nil || other == nil {
		return other == c
	}
	return c[0] == other[0] && c[1] == other[1] && c[2] == other[2]
}

var (
	zeroColor = Color{0, 0, 0}
	Red       = Color{255, 0, 0}
	Green     = Color{0, 255, 0}
	Blue      = Color{0, 0, 255}
	White     = Color{255, 255, 255}
)

type HasColor interface {
	colorAt(d time.Duration) *Color
}

func (c *Color) colorAt(time.Duration) *Color {
	return c
}

// ColorConfig is a method that takes a duration and returns a Color.
type ColorConfig func(time.Duration) *Color

func (cc ColorConfig) colorAt(d time.Duration) *Color {
	return cc(d)
}

// HSLColorConfig is a method that takes a duration and returns a HSLColor.
type HSLColorConfig func(time.Duration) *HSLColor

func (cc HSLColorConfig) colorAt(d time.Duration) *Color {
	hsl := cc(d)
	return hsl.colorAt(d)
}

// HSLColor is a triple of h,s,l: in expected ranges 0-1.
type HSLColor [3]float64

func (c *HSLColor) colorAt(time.Duration) *Color {
	if c == nil {
		return nil
	}
	rc := *c

	var r, g, b float64
	h, s, l := rc[0], rc[1], rc[2]

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)
	}

	return &Color{byte(r * 255), byte(g * 255), byte(b * 255)}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2/3-t)*6
	}
	return p
}
