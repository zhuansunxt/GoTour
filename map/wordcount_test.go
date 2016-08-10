package wordcount

import (
	"testing"
	"reflect"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		input string
		expected map[string]int
	}{
		{
			"I have a dream",
			map[string]int {
				"I" : 1,
				"have" : 1,
				"a" : 1,
				"dream" : 1,
			},
		},
		{
			"HA HA HA",
			map[string]int {
				"HA": 3,
			},
		},
		{
			"",
			map[string]int{},
		},
	}

	for _, c := range cases {
		got := WordCount(c.input)

		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("wordcount(%q) result is not correct", c.input)
			t.Errorf("Expected: %v", c.expected)
			t.Errorf("What you output: %v", got)
		}
	}
}