package fibonacci

import (
	"testing"
	"reflect"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		length int
		result []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 1}},
		{10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, c := range cases {
		length := c.length
		f := fibonacci()
		arr := []int{}
		for i := 0; i < length; i++ {
			arr = append(arr, f())
		}
		if !(reflect.DeepEqual(arr, c.result)) {
				t.Errorf("Fibonacci list of size %d testcase fails!", length)
			}
	}
}