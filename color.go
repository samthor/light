package light

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
