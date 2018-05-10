package light

import (
	"testing"
	"time"
)

func TestHSLColor(t *testing.T) {
	h := HSLColor{0, 0, 0}
	c := h.colorAt(time.Second)
	if c[0] != 0 || c[1] != 0 || c[2] != 0 {
		t.Errorf("expected black, was: %+v", c)
	}

	h = HSLColor{0, 0, 0.75}
	c = h.colorAt(time.Second)
	if c[0] != 191 || c[1] != 191 || c[2] != 191 {
		t.Errorf("expected grey, was: %+v", c)
	}

	h = HSLColor{0.9, 0.12, 0.2}
	c = h.colorAt(time.Second)
	if c[0] != 57 || c[1] != 44 || c[2] != 3 {
		t.Errorf("expected grey, was: %+v", c)
	}
}
