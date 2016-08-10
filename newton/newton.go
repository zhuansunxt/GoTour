// For computing the square root of a given float number using newton method
// The purpose of this exercise is getting familiar with golang's control flow

package newton

import (
	"math"
)

func Sqrt(x float64) (root float64) {
	init_guess := float64(1)
	conv_threshold := float64(0.001)
	delta := float64(1<<31)
	root = 0
	for ; delta > conv_threshold ; {
		root = init_guess - (math.Pow(init_guess, 2) - x) / (2 * init_guess)
		delta = math.Abs(root - init_guess)
		init_guess = root
	}
	return ;
}