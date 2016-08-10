package newton

import (
	//"fmt"
	"testing"
	"math"
)

func TestNewton(t *testing.T) {
	cases := [] struct {
		in, expected float64
	}{
		{4, 2},
		{9, 3},
		{1, 1},
		{25, 5},
	}

	threshold := float64(0.001)

	for _, c := range cases {
		got := Sqrt(c.in)
		if math.Abs(got-c.expected) > threshold {
			t.Errorf("Sqrt(%f) => %f, want %f", c.in, got, c.expected)
		}
	}
}