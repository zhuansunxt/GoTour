package reader

import (
	"testing"
	"strings"
	"bytes"
)

func TestRot13Reader(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{
			"Lbh penpxrq gur pbqr!", 
			"You cracked the code!",
		},
	}

	for _, c := range cases {
		s := strings.NewReader(c.input)
		reader := rot13reader{s}
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		got := buf.String()
		if got != c.expected {
			t.Errorf("Rot13Reader on %s not correct", c.input)
			t.Errorf("Expected : %s", c.expected)
			t.Errorf("What you got : %s", got)
		}
	}
}